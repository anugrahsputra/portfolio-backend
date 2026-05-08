package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type LanguageHandler struct {
	usecase usecase.LanguageUsecase
}

func NewLanguageHandler(u usecase.LanguageUsecase) *LanguageHandler {
	return &LanguageHandler{usecase: u}
}

func (h *LanguageHandler) CreateLanguage(c fiber.Ctx) error {
	ctx := c.Context()

	var req dto.LanguageReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToLanguageInput(&req)
	language, err := h.usecase.CreateLanguage(ctx, input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := dto.ToLanguageDTO(&language)
	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *LanguageHandler) GetLanguages(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	languages, err := h.usecase.GetLanguages(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
	}

	res := make([]dto.LanguageResp, 0, len(languages))
	for _, l := range languages {
		item := dto.ToLanguageDTO(&l)
		res = append(res, item)
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *LanguageHandler) UpdateLanguage(c fiber.Ctx) error {
	ctx := c.Context()
	languageID := c.Params("language_id")

	var req dto.LanguageUpdateReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToLanguageUpdateInput(&req)

	if err := h.usecase.UpdateLanguage(ctx, languageID, input); err != nil {
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

func (h *LanguageHandler) DeleteLanguage(c fiber.Ctx) error {
	ctx := c.Context()
	languageID := c.Params("language_id")

	if err := h.usecase.DeleteLanguage(ctx, languageID); err != nil {
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

