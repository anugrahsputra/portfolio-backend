package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	repo domain.LanguageRepository
}

func NewLanguageHandler(r domain.LanguageRepository) *LanguageHandler {
	return &LanguageHandler{repo: r}
}

func (h *LanguageHandler) CreateLanguage(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.LanguageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToLanguageInput(&req)
	language, err := h.repo.CreateLanguage(ctx, input)
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

	languages, err := h.repo.GetLanguages(ctx, profileID)
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

	if err := h.repo.UpdateLanguage(ctx, languageID, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

func (h *LanguageHandler) DeleteLanguage(c *gin.Context) {
	ctx := c.Request.Context()
	languageID := c.Param("language_id")

	if err := h.repo.DeleteLanguage(ctx, languageID); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}
