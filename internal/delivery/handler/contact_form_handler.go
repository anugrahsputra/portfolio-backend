package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
)

type ContactFormHandler struct {
	usecase usecase.EmailContactUsecase
}

func NewContactFormHandler(u usecase.EmailContactUsecase) *ContactFormHandler {
	return &ContactFormHandler{usecase: u}
}

func (h *ContactFormHandler) SendMail(w http.ResponseWriter, r *http.Request) {
	var req dto.ContactFormReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToContactFormInput(&req)
	if err := h.usecase.SendEmail(r.Context(), input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "email submitted")
}
