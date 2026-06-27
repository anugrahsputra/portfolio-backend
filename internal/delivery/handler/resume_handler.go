package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	repo domain.ResumeRepository
}

func NewResumeHandler(r domain.ResumeRepository) *ResumeHandler {
	return &ResumeHandler{repo: r}
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_id")

	resume, err := h.repo.GetResume(ctx, id)
	if err != nil {
		ResponseError(c, http.StatusNotFound, "Resume not found")
		return
	}

	res := dto.ToResumeDTO(resume)
	ResponseJSON(c, http.StatusOK, "Success get resume", res)
}
