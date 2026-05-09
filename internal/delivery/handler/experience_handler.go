package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	usecase usecase.ExperienceUsecase
}

func NewExperienceHandler(u usecase.ExperienceUsecase) *ExperienceHandler {
	return &ExperienceHandler{usecase: u}
}

func (h *ExperienceHandler) CreateExperience(c *gin.Context) {
	var req dto.ExperienceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToExperienceInput(&req)
	experience, err := h.usecase.CreateExperience(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	res := dto.ToExperienceDTO(&experience)
	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) GetExperiences(c *gin.Context) {
	profileID := c.Param("profile_id")

	experiences, err := h.usecase.GetExperiences(c.Request.Context(), profileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "bad request",
		})
		return
	}

	res := make([]dto.ExperienceResp, 0, len(experiences))
	for _, ex := range experiences {
		item := dto.ToExperienceDTO(&ex)
		res = append(res, item)
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) UpdateExperience(c *gin.Context) {
	expID := c.Param("experience_id")

	var req dto.ExperienceUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToExperienceUpdateInput(&req)
	experience, err := h.usecase.UpdateExperience(c.Request.Context(), expID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	res := dto.ToExperienceDTO(&experience)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) DeleteExperience(c *gin.Context) {
	expID := c.Param("experience_id")

	if err := h.usecase.DeleteExperience(c.Request.Context(), expID); err != nil {
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
