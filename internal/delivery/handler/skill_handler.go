package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	repo domain.SkillRepository
}

func NewSkillHandler(r domain.SkillRepository) *SkillHandler {
	return &SkillHandler{repo: r}
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

	skill, err := h.repo.CreateSkill(ctx, input)
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

	skill, err := h.repo.GetSkills(ctx, profileID)
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

	if err := h.repo.UpdateSkill(ctx, skillId, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	ctx := c.Request.Context()
	skillId := c.Param("skill_id")

	if err := h.repo.DeleteSkill(ctx, skillId); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}
