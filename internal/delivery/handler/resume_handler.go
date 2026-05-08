package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ResumeHandler struct {
	usecase usecase.ResumeUsecase
}

func NewResumeHandler(u usecase.ResumeUsecase) *ResumeHandler {
	return &ResumeHandler{usecase: u}
}

func (h *ResumeHandler) GetResume(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profile_id")

	resume, err := h.usecase.GetResume(r.Context(), id)
	if err != nil {
		ResponseError(w, r, http.StatusNotFound, "Resume not found")
		return
	}

	res := dto.ToResumeDTO(resume)
	ResponseJSON(w, r, http.StatusOK, "Success get resume", res)
}
