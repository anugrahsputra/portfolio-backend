package handler

import (
	"fmt"
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
	ctx := c.Request.Context()

	var req dto.EducationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := dto.ToEducationInput(&req)
	if err := h.usecase.CreateEducation(ctx, input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.NoDataResponse{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *EducationHandler) GetEducation(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	educations, err := h.usecase.GetEducations(ctx, profileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
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
	ctx := c.Request.Context()
	eduID := c.Param("education_id")

	var req dto.EducationUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := dto.ToEducationUpdateInput(&req)
	if err := h.usecase.UpdateEducation(ctx, eduID, input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *EducationHandler) DeleteEducation(c *gin.Context) {
	ctx := c.Request.Context()
	eduID := c.Param("education_id")

	if err := h.usecase.DeleteEducation(ctx, eduID); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
