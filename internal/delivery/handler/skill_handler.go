package handler

import (
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
		ResponseError(c, http.StatusBadRequest, "invalid request body")
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
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToSkillDTO(&skill)
	ResponseJSON(c, http.StatusCreated, "success", res)
}

func (h *SkillHandler) GetSkills(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	skill, err := h.usecase.GetSkills(ctx, profileID)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToSkillDTO(&skill)
	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	ctx := c.Request.Context()
	skillId := c.Param("skill_id")

	var req dto.SkillUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.SkillUpdateInput{
		Tools:        req.Tools,
		Technologies: req.Technologies,
		HardSkills:   req.HardSkills,
		SoftSkills:   req.SoftSkills,
	}

	if err := h.usecase.UpdateSkill(ctx, skillId, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	ctx := c.Request.Context()
	skillId := c.Param("skill_id")

	if err := h.usecase.DeleteSkill(ctx, skillId); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

