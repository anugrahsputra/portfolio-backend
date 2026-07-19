package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	uc domain.ResumeUsecase
}

func NewResumeHandler(u domain.ResumeUsecase) *ResumeHandler {
	return &ResumeHandler{uc: u}
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_id")

	resume, err := h.uc.GetResume(ctx, id)
	if err != nil {
		ResponseError(c, http.StatusNotFound, "Resume not found")
		return
	}

	res := dto.ToResumeDTO(resume)
	ResponseJSON(c, http.StatusOK, "Success get resume", res)
}
