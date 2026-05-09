package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type EducationHandler struct {
	usecase usecase.EducationUsecase
}

func NewEducationHandler(u usecase.EducationUsecase) *EducationHandler {
	return &EducationHandler{usecase: u}
}

func (h *EducationHandler) CreateEducation(c *gin.Context) {
	var req dto.EducationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToEducationInput(&req)
	if err := h.usecase.CreateEducation(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.NoDataResponse{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *EducationHandler) GetEducation(c *gin.Context) {
	profileID := c.Param("profile_id")

	educations, err := h.usecase.GetEducations(c.Request.Context(), profileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "bad request",
		})
		return
	}

	res := make([]dto.EducationResp, 0, len(educations))
	for _, ed := range educations {
		item := dto.ToEducationDTO(&ed)
		res = append(res, item)
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *EducationHandler) UpdateEducation(c *gin.Context) {
	eduID := c.Param("education_id")

	var req dto.EducationUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToEducationUpdateInput(&req)
	if err := h.usecase.UpdateEducation(c.Request.Context(), eduID, input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *EducationHandler) DeleteEducation(c *gin.Context) {
	eduID := c.Param("education_id")

	if err := h.usecase.DeleteEducation(c.Request.Context(), eduID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
