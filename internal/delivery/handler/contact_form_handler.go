package handler

import (
	"fmt"
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
	ctx := c.Request.Context()

	var req dto.ContactFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := dto.ToContactFormInput(&req)
	if err := h.usecase.SendEmail(ctx, input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "email submitted",
	})
}
