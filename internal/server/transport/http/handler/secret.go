package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/timac11/goph-keeper/internal/common/model"
	"github.com/timac11/goph-keeper/internal/server/errors"
)

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

func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()

	secret, err := h.service.GetSecret(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnBody, err := json.Marshal(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnBody)
}

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

func (h *Handler) GetSecretList(w http.ResponseWriter, r *http.Request) {
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
	err := r.ParseMultipartForm(32 << 20) // TODO: add max file size
	if err != nil {
		return nil, errors.SecretInvalidForm
	}

	name := r.FormValue("name")
	uploadType := r.FormValue("type")
	metadata := r.FormValue("metadata")

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

	dataBytes, err := io.ReadAll(dataFile)
	if err != nil {
		return nil, errors.SecretFileFailedToRead
	}

	pubFile, _, err := r.FormFile("publicKey")
	if err != nil {
		return nil, errors.PublicKeyIsRequired
	}
	defer pubFile.Close()

	publicKeyBytes, err := io.ReadAll(pubFile)
	if err != nil {
		return nil, errors.PublicKeyFailedToRead
	}

	return &model.SecretCreateDto{Name: name,
		Type:      uploadType,
		Metadata:  metadata,
		Data:      dataBytes,
		PublicKey: publicKeyBytes,
	}, nil
}
