package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type ContactFormHandler struct {
	repo domain.EmailContactRepository
}

func NewContactFormHandler(r domain.EmailContactRepository) *ContactFormHandler {
	return &ContactFormHandler{repo: r}
}

func (h *ContactFormHandler) SendMail(c *gin.Context) {
	var req dto.ContactFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToContactFormInput(&req)

	if err := h.repo.SendEmail(c.Request.Context(), input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "email submitted")
}


