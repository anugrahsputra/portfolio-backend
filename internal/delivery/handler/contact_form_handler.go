package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type ContactFormHandler struct {
	uc domain.ContactFormUsecase
}

func NewContactFormHandler(u domain.ContactFormUsecase) *ContactFormHandler {
	return &ContactFormHandler{uc: u}
}

func (h *ContactFormHandler) SendMail(c *gin.Context) {
	var req dto.ContactFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToContactFormInput(&req)

	if err := h.uc.SendEmail(c.Request.Context(), input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "email submitted")
}


