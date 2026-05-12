package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ContactFormHandler struct {
	usecase usecase.EmailContactUsecase
}

func NewContactFormHandler(u usecase.EmailContactUsecase) *ContactFormHandler {
	return &ContactFormHandler{usecase: u}
}

func (h *ContactFormHandler) SendMail(c *gin.Context) {
	var req dto.ContactFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToContactFormInput(&req)
	if err := h.usecase.SendEmail(c.Request.Context(), input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "email submitted")
}
