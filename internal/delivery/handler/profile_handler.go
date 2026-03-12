package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(u usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecase: u}
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.ProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := domain.ProfileInput{
		Name:    req.Name,
		About:   req.About,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	}

	res, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "failed to create profile",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Success create profile",
		Data:    res,
	})
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	profile, err := h.usecase.GetProfile(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "failed to get profile",
		})
		return
	}

	res := dto.ToProfileDTO(profile)

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success get profile",
		Data:    res,
	})
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	var req dto.ProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := domain.ProfileUpdateInput{
		Name:    &req.Name,
		About:   &req.About,
		Address: &req.Address,
		Email:   &req.Email,
		Phone:   &req.Phone,
	}

	if err := h.usecase.UpdateProfile(ctx, id, input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "failed to update profile",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success update profile",
	})
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	if err := h.usecase.DeleteProfile(ctx, id); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "failed to delete profile",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success delete profile",
	})
}
