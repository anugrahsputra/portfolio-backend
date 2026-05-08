package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/go-chi/render"
)

func ResponseJSON(w http.ResponseWriter, r *http.Request, status int, message string, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, dto.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, status int, message string) {
	render.Status(r, status)
	render.JSON(w, r, dto.NoDataResponse{
		Status:  status,
		Message: message,
	})
}
