package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	usecase usecase.SkillUsecase
}

func NewSkillHandler(u usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{usecase: u}
}

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.SkillReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
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
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed create skill: %v", err),
		})
		return
	}

	res := dto.ToSkillDTO(&skill)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *SkillHandler) GetSkills(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	skill, err := h.usecase.GetSkills(ctx, profileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	res := dto.ToSkillDTO(&skill)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	ctx := c.Request.Context()
	skillId := c.Param("skill_id")

	var req dto.SkillUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := domain.SkillUpdateInput{
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	if err := h.usecase.UpdateSkill(ctx, skillId, input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update skill: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	ctx := c.Request.Context()
	skillId := c.Param("skill_id")

	if err := h.usecase.DeleteSkill(ctx, skillId); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to delete skill: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
