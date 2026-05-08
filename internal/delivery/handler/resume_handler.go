package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ResumeHandler struct {
	usecase usecase.ResumeUsecase
}

func NewResumeHandler(u usecase.ResumeUsecase) *ResumeHandler {
	return &ResumeHandler{usecase: u}
}

func (h *ResumeHandler) GetResume(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("profile_id")

	resume, err := h.usecase.GetResume(ctx, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.NoDataResponse{
			Status:  http.StatusNotFound,
			Message: "Resume not found",
		})
	}

	res := dto.ToResumeDTO(resume)

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "Success get resume",
		Data:    res,
	})
}

