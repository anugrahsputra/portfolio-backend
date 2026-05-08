package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ExperienceHandler struct {
	usecase usecase.ExperienceUsecase
}

func NewExperienceHandler(u usecase.ExperienceUsecase) *ExperienceHandler {
	return &ExperienceHandler{usecase: u}
}

func (h *ExperienceHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	var req dto.ExperienceReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToExperienceInput(&req)
	experience, err := h.usecase.CreateExperience(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToExperienceDTO(&experience)
	ResponseJSON(w, r, http.StatusCreated, "success", res)
}

func (h *ExperienceHandler) GetExperiences(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	experiences, err := h.usecase.GetExperiences(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.ExperienceResp, 0, len(experiences))
	for _, ex := range experiences {
		item := dto.ToExperienceDTO(&ex)
		res = append(res, item)
	}

	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *ExperienceHandler) UpdateExperience(w http.ResponseWriter, r *http.Request) {
	expID := chi.URLParam(r, "experience_id")

	var req dto.ExperienceUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToExperienceUpdateInput(&req)
	experience, err := h.usecase.UpdateExperience(r.Context(), expID, input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToExperienceDTO(&experience)
	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *ExperienceHandler) DeleteExperience(w http.ResponseWriter, r *http.Request) {
	expID := chi.URLParam(r, "experience_id")

	if err := h.usecase.DeleteExperience(r.Context(), expID); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}
