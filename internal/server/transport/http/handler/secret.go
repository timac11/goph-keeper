package handler

import (
	"net/http"
)

func (h *Handler) CreateSecret(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) DeleteSecret(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) UpdateSecret(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetSecretList(w http.ResponseWriter, r *http.Request) {}
