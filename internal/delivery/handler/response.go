package handler

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, dto.NoDataResponse{
		Status:  status,
		Message: message,
	})
}
