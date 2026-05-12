package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	usecase usecase.ResumeUsecase
}

func NewResumeHandler(u usecase.ResumeUsecase) *ResumeHandler {
	return &ResumeHandler{usecase: u}
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_id")

	resume, err := h.usecase.GetResume(ctx, id)
	if err != nil {
		ResponseError(c, http.StatusNotFound, "Resume not found")
		return
	}

	res := dto.ToResumeDTO(resume)
	ResponseJSON(c, http.StatusOK, "Success get resume", res)
}

