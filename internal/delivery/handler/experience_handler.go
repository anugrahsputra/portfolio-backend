package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	repo domain.ExperienceRepository
}

func NewExperienceHandler(r domain.ExperienceRepository) *ExperienceHandler {
	return &ExperienceHandler{repo: r}
}

func (h *ExperienceHandler) CreateExperience(c *gin.Context) {
	var req dto.ExperienceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToExperienceInput(&req)
	experience, err := h.repo.CreateExperience(c.Request.Context(), input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToExperienceDTO(&experience)
	ResponseJSON(c, http.StatusCreated, "success", res)
}

func (h *ExperienceHandler) GetExperiences(c *gin.Context) {
	profileID := c.Param("profile_id")

	experiences, err := h.repo.GetExperiences(c.Request.Context(), profileID)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.ExperienceResp, 0, len(experiences))
	for _, ex := range experiences {
		item := dto.ToExperienceDTO(&ex)
		res = append(res, item)
	}

	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *ExperienceHandler) UpdateExperience(c *gin.Context) {
	expID := c.Param("experience_id")

	var req dto.ExperienceUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToExperienceUpdateInput(&req)
	experience, err := h.repo.UpdateExperience(c.Request.Context(), expID, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToExperienceDTO(&experience)
	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *ExperienceHandler) DeleteExperience(c *gin.Context) {
	expID := c.Param("experience_id")

	if err := h.repo.DeleteExperience(c.Request.Context(), expID); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}
