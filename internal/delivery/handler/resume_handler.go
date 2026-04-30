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
	id := c.Param("id")

	resume, err := h.usecase.GetResume(ctx, id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Resume not found",
		})
		return
	}

	res := dto.ToResumeDTO(resume)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get resume",
		"data":    res,
	})
}
