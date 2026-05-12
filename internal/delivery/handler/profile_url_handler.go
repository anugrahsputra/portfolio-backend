package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProfileUrlHandler struct {
	usecase usecase.ProfileUrlUsecase
}

func NewProfileUrlHandler(u usecase.ProfileUrlUsecase) *ProfileUrlHandler {
	return &ProfileUrlHandler{usecase: u}
}

func (h *ProfileUrlHandler) CreateProfileUrl(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.ProfileUrlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUrlInput{
		ProfileID: req.ProfileID,
		Label:     req.Label,
		Url:       req.Url,
	}

	profileUrl, err := h.usecase.CreateProfileUrl(ctx, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToProfileUrlDTO(profileUrl)
	ResponseJSON(c, http.StatusCreated, "created", res)
}

func (h *ProfileUrlHandler) GetProfileURL(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	profileUrls, err := h.usecase.GetProfileUrl(ctx, profileID)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := make([]dto.ProfileUrlResp, 0, len(profileUrls))
	for _, pUrl := range profileUrls {
		item := dto.ToProfileUrlDTO(&pUrl)
		res = append(res, item)
	}

	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *ProfileUrlHandler) GetProfileUrlByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	profileUrl, err := h.usecase.GetProfileUrlByID(ctx, id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToProfileUrlDTO(&profileUrl)
	ResponseJSON(c, http.StatusOK, "success", res)
}

func (h *ProfileUrlHandler) UpdateProfileUrl(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	var req dto.ProfileUrlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUrlUpdateInput{
		ProfileID: &req.ProfileID,
		Label:     &req.Label,
		Url:       &req.Url,
	}

	if err := h.usecase.UpdateProfileUrl(ctx, id, input); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "Profile url updated")
}

func (h *ProfileUrlHandler) DeleteProfileUrl(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	if err := h.usecase.DeleteProfileUrl(ctx, id); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "Profile url deleted")
}

