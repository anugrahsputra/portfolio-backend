package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ContactFormHandler struct {
	usecase usecase.EmailContactUsecase
}

func NewContactFormHandler(u usecase.EmailContactUsecase) *ContactFormHandler {
	return &ContactFormHandler{usecase: u}
}

func (h *ContactFormHandler) SendMail(c fiber.Ctx) error {
	ctx := c.Context()
	var req dto.ContactFormReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := dto.ToContactFormInput(&req)
	if err := h.usecase.SendEmail(ctx, input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("failed to send email: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "email submitted",
	})
}

