package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ExperienceHandler struct {
	usecase usecase.ExperienceUsecase
}

func NewExperienceHandler(u usecase.ExperienceUsecase) *ExperienceHandler {
	return &ExperienceHandler{usecase: u}
}

func (h *ExperienceHandler) CreateExperience(c fiber.Ctx) error {
	ctx := c.Context()

	var req dto.ExperienceReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToExperienceInput(&req)
	experience, err := h.usecase.CreateExperience(ctx, input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := dto.ToExperienceDTO(&experience)
	return c.Status(http.StatusCreated).JSON(dto.Response{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) GetExperiences(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	experiences, err := h.usecase.GetExperiences(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := make([]dto.ExperienceResp, 0, len(experiences))
	for _, ex := range experiences {
		item := dto.ToExperienceDTO(&ex)
		res = append(res, item)
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) UpdateExperience(c fiber.Ctx) error {
	ctx := c.Context()
	expID := c.Params("experience_id")

	var req dto.ExperienceUpdateReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToExperienceUpdateInput(&req)
	experience, err := h.usecase.UpdateExperience(ctx, expID, input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := dto.ToExperienceDTO(&experience)
	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ExperienceHandler) DeleteExperience(c fiber.Ctx) error {
	ctx := c.Context()
	expID := c.Params("experience_id")

	if err := h.usecase.DeleteExperience(ctx, expID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

