package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ProfileUrlHandler struct {
	usecase usecase.ProfileUrlUsecase
}

func NewProfileUrlHandler(u usecase.ProfileUrlUsecase) *ProfileUrlHandler {
	return &ProfileUrlHandler{usecase: u}
}

func (h *ProfileUrlHandler) CreateProfileUrl(c fiber.Ctx) error {
	ctx := c.Context()

	var req dto.ProfileUrlReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
	}

	input := domain.ProfileUrlInput{
		ProfileID: req.ProfileID,
		Label:     req.Label,
		Url:       req.Url,
	}

	profileUrl, err := h.usecase.CreateProfileUrl(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to create profile url: %v", err),
		})
	}

	res := dto.ToProfileUrlDTO(profileUrl)

	return c.Status(http.StatusCreated).JSON(dto.Response{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) GetProfileURL(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	profileUrls, err := h.usecase.GetProfileUrl(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to get profile url: %v", err),
		})
	}

	res := make([]dto.ProfileUrlResp, 0, len(profileUrls))
	for _, pUrl := range profileUrls {
		item := dto.ToProfileUrlDTO(&pUrl)
		res = append(res, item)
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) GetProfileUrlByID(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("profile_url_id")

	profileUrl, err := h.usecase.GetProfileUrlByID(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to get profile url: %v", err),
		})
	}

	res := dto.ToProfileUrlDTO(&profileUrl)

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) UpdateProfileUrl(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("profile_url_id")

	var req dto.ProfileUrlReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
	}

	input := domain.ProfileUrlUpdateInput{
		ProfileID: &req.ProfileID,
		Label:     &req.Label,
		Url:       &req.Url,
	}

	if err := h.usecase.UpdateProfileUrl(ctx, id, input); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to update profile url: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "Profile url updated",
	})
}

func (h *ProfileUrlHandler) DeleteProfileUrl(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("profile_url_id")

	if err := h.usecase.DeleteProfileUrl(ctx, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to delete profile url: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "Profile url deleted",
	})
}

