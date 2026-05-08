package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type EducationHandler struct {
	usecase usecase.EducationUsecase
}

func NewEducationHandler(u usecase.EducationUsecase) *EducationHandler {
	return &EducationHandler{usecase: u}
}

func (h *EducationHandler) CreateEducation(w http.ResponseWriter, r *http.Request) {
	var req dto.EducationReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToEducationInput(&req)
	if err := h.usecase.CreateEducation(r.Context(), input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusCreated, "success")
}

func (h *EducationHandler) GetEducation(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	educations, err := h.usecase.GetEducations(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.EducationResp, 0, len(educations))
	for _, ed := range educations {
		item := dto.ToEducationDTO(&ed)
		res = append(res, item)
	}

	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *EducationHandler) UpdateEducation(w http.ResponseWriter, r *http.Request) {
	eduID := chi.URLParam(r, "education_id")

	var req dto.EducationUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToEducationUpdateInput(&req)
	if err := h.usecase.UpdateEducation(r.Context(), eduID, input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}

func (h *EducationHandler) DeleteEducation(w http.ResponseWriter, r *http.Request) {
	eduID := chi.URLParam(r, "education_id")

	if err := h.usecase.DeleteEducation(r.Context(), eduID); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}
