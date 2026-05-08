package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(u usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecase: u}
}

func (h *ProfileHandler) CreateProfile(c fiber.Ctx) error {
	ctx := c.Context()
	var req dto.ProfileReq

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
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
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create profile",
		})
	}

	res := dto.ToProfileDTO(profile)
	return c.Status(http.StatusCreated).JSON(dto.Response{
		Status:  http.StatusCreated,
		Message: "Success create profile",
		Data:    res,
	})
}

func (h *ProfileHandler) GetProfile(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	profile, err := h.usecase.GetProfile(ctx, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.NoDataResponse{
			Status:  http.StatusNotFound,
			Message: "Profile not found",
		})
	}

	res := dto.ToProfilePublicDTO(profile)

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success get profile",
		Data:    res,
	})
}

func (h *ProfileHandler) UpdateProfile(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.ProfileReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
	}

	input := domain.ProfileUpdateInput{
		Name:    &req.Name,
		About:   &req.About,
		Address: &req.Address,
		Email:   &req.Email,
		Phone:   &req.Phone,
	}

	if err := h.usecase.UpdateProfile(ctx, id, input); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update profile",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success update profile",
	})
}

func (h *ProfileHandler) DeleteProfile(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.usecase.DeleteProfile(ctx, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete profile",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success delete profile",
	})
}

