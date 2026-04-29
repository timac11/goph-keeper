package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/timac11/goph-keeper/internal/common/logger"
	"github.com/timac11/goph-keeper/internal/common/model"
	"github.com/timac11/goph-keeper/internal/server/errors"
)

// CreateSecret godoc
// @Summary Upload secret with metadata
// @Description Secret multipart/form-data with fields:
// @Description - name: string
// @Description - metadata: string
// @Description - data: file (binary)
// @Description - type: one of [FILE, CARD, CREDS]
// @Description - publicKey: file (binary)
// @Tags Secret
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param data formData file true "Data file (binary)"
// @Param type formData string true "Type (file|card|creds)"
// @Param publicKey formData file true "Public key (binary)"
// @Success 200 {object} model.SecretInfoDto
// @Router /secret [post]
// @Security     ApiKeyAuth
func (h *Handler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	secret, err := h.extractSecretModelFromForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	log := logger.LoggerFromContext(r.Context())

	_, err = h.service.CreateSecret(ctx, secret)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetSecret godoc
// @Summary		Get secret
// @Description	Returns secret fields and secret data file as multipart/form-data.
// @Tags Secret
// @Produce		multipart/form-data
// @Param			id	path	string	true	"Secret id"
// @Success		200	{string}	string	"multipart/form-data response with fields: id, name, metadata, publicKey, type and file field data"
// @Failure		500	{string}	string	"Internal server error"
// @Router       /secret/{id} [get]
// @Security     ApiKeyAuth
func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()

	secret, err := h.service.GetSecret(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.Open(secret.DataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("X-Secret-Name", secret.Name)
	w.Header().Set("X-Secret-Type", secret.Type)
	w.Header().Set("X-Secret-Public-Key", secret.PublicKey)
	w.Header().Set("X-Secret-Metadata", secret.Metadata)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, secret.Name))
	w.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteSecret godoc
// @Summary      Remove Secret
// @Description  Remove secret by id
// @Tags Secret
// @Param        id   path    string  true  "Secret id"
// @Router       /secret/{id} [delete]
// @Security     ApiKeyAuth
func (h *Handler) DeleteSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()

	err := h.service.DeleteSecret(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetSecret godoc
// @Summary      Return Secrets list
// @Description  Get secrets list
// @Tags Secret
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.SecretInfoDto
// @Router       /secret [get]
// @Security     ApiKeyAuth
func (h *Handler) GetSecretsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	secrets, err := h.service.GetSecretsList(ctx)

	returnBody, err := json.Marshal(secrets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(returnBody))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnBody)
}

func (h *Handler) extractSecretModelFromForm(r *http.Request) (*model.SecretCreateDto, error) {
	log := logger.LoggerFromContext(r.Context())

	name := r.Header.Get("X-Secret-Name")
	secretType := r.Header.Get("X-Secret-Type")
	metadata := r.Header.Get("X-Secret-Metadata")
	publicKey := r.Header.Get("X-Secret-Public-Key")

	log.Info("parsed fields",
		zap.String("name", name),
		zap.String("secretType", secretType),
		zap.String("publicKey", publicKey),
	)

	if name == "" {
		return nil, errors.SecretNameIsRequired
	}

	if secretType == "" {
		return nil, errors.SecretInvalidType
	}

	if publicKey == "" {
		return nil, errors.PublicKeyIsRequired
	}

	fileName := filepath.Base(name)
	dataPath := filepath.Join("./uploads", fileName)

	dst, err := os.Create(dataPath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer dst.Close()

	// move logic to service
	if _, err := io.Copy(dst, r.Body); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &model.SecretCreateDto{
		Name:      name,
		Type:      secretType,
		Metadata:  metadata,
		PublicKey: publicKey,
		DataPath:  dataPath,
	}, nil
}
