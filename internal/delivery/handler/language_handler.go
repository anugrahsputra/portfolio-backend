package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	usecase usecase.LanguageUsecase
}

func NewLanguageHandler(u usecase.LanguageUsecase) *LanguageHandler {
	return &LanguageHandler{usecase: u}
}

func (h *LanguageHandler) CreateLanguage(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.LanguageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := dto.ToLanguageInput(&req)
	language, err := h.usecase.CreateLanguage(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	res := dto.ToLanguageDTO(&language)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *LanguageHandler) GetLanguages(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	languages, err := h.usecase.GetLanguages(ctx, profileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	res := make([]dto.LanguageResp, 0, len(languages))
	for _, l := range languages {
		item := dto.ToLanguageDTO(&l)
		res = append(res, item)
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *LanguageHandler) UpdateLanguage(c *gin.Context) {
	ctx := c.Request.Context()
	languageID := c.Param("language_id")

	var req dto.LanguageUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error: %v", err),
		})
		return
	}

	input := dto.ToLanguageUpdateInput(&req)

	if err := h.usecase.UpdateLanguage(ctx, languageID, input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *LanguageHandler) DeleteLanguage(c *gin.Context) {
	ctx := c.Request.Context()
	languageID := c.Param("language_id")

	if err := h.usecase.DeleteLanguage(ctx, languageID); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}
