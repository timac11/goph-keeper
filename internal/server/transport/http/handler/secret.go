package handler

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gorilla/mux"

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

	_, err = h.service.CreateSecret(ctx, secret)

	if err != nil {
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

	mw := multipart.NewWriter(w)
	defer mw.Close()

	w.Header().Set("Content-Type", mw.FormDataContentType())

	_ = mw.WriteField("id", secret.ID)
	_ = mw.WriteField("name", secret.Name)
	_ = mw.WriteField("metadata", secret.Metadata)
	_ = mw.WriteField("publicKey", secret.PublicKey)
	_ = mw.WriteField("type", secret.Type)

	part, err := mw.CreateFormFile("data", secret.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.Open(secret.DataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(part, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnBody)
}

func (h *Handler) extractSecretModelFromForm(r *http.Request) (*model.SecretCreateDto, error) {
	name := r.FormValue("name")
	uploadType := r.FormValue("type")
	metadata := r.FormValue("metadata")
	publicKey := r.FormValue("publicKey")

	if name == "" {
		return nil, errors.SecretNameIsRequired
	}

	switch uploadType {
	case model.Auth, model.Card, model.File:
	default:
		return nil, errors.SecretInvalidType
	}

	dataFile, _, err := r.FormFile("data")
	if err != nil {
		return nil, errors.SecretFileIsRequired
	}
	defer dataFile.Close()

	dataPath := "./uploads/" + name

	dst, err := os.Create(dataPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, dataFile); err != nil {
		return nil, err
	}

	return &model.SecretCreateDto{
		Name:      name,
		Type:      uploadType,
		Metadata:  metadata,
		DataPath:  dataPath,
		PublicKey: publicKey,
	}, nil
}
