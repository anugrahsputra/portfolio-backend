package handler

import (
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
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToLanguageInput(&req)
	language, err := h.usecase.CreateLanguage(ctx, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToLanguageDTO(&language)
	ResponseJSON(c, http.StatusCreated, "success", res)
}

func (h *LanguageHandler) GetLanguages(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	languages, err := h.usecase.GetLanguages(ctx, profileID)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.LanguageResp, 0, len(languages))
	for _, l := range languages {
		item := dto.ToLanguageDTO(&l)
		res = append(res, item)
	}

	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *LanguageHandler) UpdateLanguage(c *gin.Context) {
	ctx := c.Request.Context()
	languageID := c.Param("language_id")

	var req dto.LanguageUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToLanguageUpdateInput(&req)

	if err := h.usecase.UpdateLanguage(ctx, languageID, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

func (h *LanguageHandler) DeleteLanguage(c *gin.Context) {
	ctx := c.Request.Context()
	languageID := c.Param("language_id")

	if err := h.usecase.DeleteLanguage(ctx, languageID); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

