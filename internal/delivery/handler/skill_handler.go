package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type SkillHandler struct {
	usecase usecase.SkillUsecase
}

func NewSkillHandler(u usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{usecase: u}
}

func (h *SkillHandler) CreateSkill(c fiber.Ctx) error {
	ctx := c.Context()

	var req dto.SkillReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := domain.SkillInput{
		ProfileID:    req.ProfileID,
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	skill, err := h.usecase.CreateSkill(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed create skill: %v", err),
		})
	}

	res := dto.ToSkillDTO(&skill)
	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *SkillHandler) GetSkills(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	skill, err := h.usecase.GetSkills(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	res := dto.ToSkillDTO(&skill)
	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *SkillHandler) UpdateSkill(c fiber.Ctx) error {
	ctx := c.Context()
	skillId := c.Params("skill_id")

	var req dto.SkillUpdateReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
	}

	input := domain.SkillUpdateInput{
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	if err := h.usecase.UpdateSkill(ctx, skillId, input); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update skill: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *SkillHandler) DeleteSkill(c fiber.Ctx) error {
	ctx := c.Context()
	skillId := c.Params("skill_id")

	if err := h.usecase.DeleteSkill(ctx, skillId); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to delete skill: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

