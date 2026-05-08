package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type SkillHandler struct {
	usecase usecase.SkillUsecase
}

func NewSkillHandler(u usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{usecase: u}
}

func (h *SkillHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	var req dto.SkillReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.SkillInput{
		ProfileID:    req.ProfileID,
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	skill, err := h.usecase.CreateSkill(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToSkillDTO(&skill)
	ResponseJSON(w, r, http.StatusCreated, "success", res)
}

func (h *SkillHandler) GetSkills(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	skill, err := h.usecase.GetSkills(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToSkillDTO(&skill)
	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *SkillHandler) UpdateSkill(w http.ResponseWriter, r *http.Request) {
	skillId := chi.URLParam(r, "skill_id")

	var req dto.SkillUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.SkillUpdateInput{
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	if err := h.usecase.UpdateSkill(r.Context(), skillId, input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}

func (h *SkillHandler) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	skillId := chi.URLParam(r, "skill_id")

	if err := h.usecase.DeleteSkill(r.Context(), skillId); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}
