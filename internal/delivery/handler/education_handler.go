package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type EducationHandler struct {
	repo domain.EducationRepository
}

func NewEducationHandler(r domain.EducationRepository) *EducationHandler {
	return &EducationHandler{repo: r}
}

func (h *EducationHandler) CreateEducation(c *gin.Context) {
	var req dto.EducationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToEducationInput(&req)
	if err := h.repo.CreateEducation(c.Request.Context(), input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusCreated, "success")
}

func (h *EducationHandler) GetEducation(c *gin.Context) {
	profileID := c.Param("profile_id")

	educations, err := h.repo.GetEducations(c.Request.Context(), profileID)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.EducationResp, 0, len(educations))
	for _, ed := range educations {
		item := dto.ToEducationDTO(&ed)
		res = append(res, item)
	}

	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *EducationHandler) UpdateEducation(c *gin.Context) {
	eduID := c.Param("education_id")

	var req dto.EducationUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToEducationUpdateInput(&req)
	if err := h.repo.UpdateEducation(c.Request.Context(), eduID, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

func (h *EducationHandler) DeleteEducation(c *gin.Context) {
	eduID := c.Param("education_id")

	if err := h.repo.DeleteEducation(c.Request.Context(), eduID); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}
