package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type EducationHandler struct {
	usecase usecase.EducationUsecase
}

func NewEducationHandler(u usecase.EducationUsecase) *EducationHandler {
	return &EducationHandler{usecase: u}
}

func (h *EducationHandler) CreateEducation(c fiber.Ctx) error {
	ctx := c.Context()
	var req dto.EducationReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToEducationInput(&req)
	if err := h.usecase.CreateEducation(ctx, input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.NoDataResponse{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *EducationHandler) GetEducation(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	educations, err := h.usecase.GetEducations(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := make([]dto.EducationResp, 0, len(educations))
	for _, ed := range educations {
		item := dto.ToEducationDTO(&ed)
		res = append(res, item)
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *EducationHandler) UpdateEducation(c fiber.Ctx) error {
	ctx := c.Context()
	eduID := c.Params("education_id")

	var req dto.EducationUpdateReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToEducationUpdateInput(&req)
	if err := h.usecase.UpdateEducation(ctx, eduID, input); err != nil {
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

func (h *EducationHandler) DeleteEducation(c fiber.Ctx) error {
	ctx := c.Context()
	eduID := c.Params("education_id")

	if err := h.usecase.DeleteEducation(ctx, eduID); err != nil {
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

