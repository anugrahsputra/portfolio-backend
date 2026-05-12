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
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileInput{
		Name:    req.Name,
		Title:   req.Title,
		About:   req.About,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	}

	profile, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "Failed to create profile")
		return
	}

	res := dto.ToProfileDTO(profile)
	ResponseJSON(c, http.StatusCreated, "Success create profile", res)
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	profile, err := h.usecase.GetProfile(ctx, id)
	if err != nil {
		ResponseError(c, http.StatusNotFound, "Profile not found")
		return
	}

	res := dto.ToProfilePublicDTO(profile)
	ResponseJSON(c, http.StatusOK, "success get profile", res)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var req dto.ProfileUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUpdateInput{
		Name:    req.Name,
		Title:   req.Title,
		About:   req.About,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	}

	if err := h.usecase.UpdateProfile(ctx, id, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	ResponseError(c, http.StatusOK, "success update profile")
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.usecase.DeleteProfile(ctx, id); err != nil {
		ResponseError(c, http.StatusInternalServerError, "Failed to delete profile")
		return
	}

	ResponseError(c, http.StatusOK, "success delete profile")
}



