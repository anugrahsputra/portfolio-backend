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
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := domain.ProfileUrlInput{
		ProfileID: req.ProfileID,
		Label:     req.Label,
		Url:       req.Url,
	}

	profileUrl, err := h.usecase.CreateProfileUrl(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	res := dto.ToProfileUrlDTO(profileUrl)
	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) GetProfileURL(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	profileUrls, err := h.usecase.GetProfileUrl(ctx, profileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	res := make([]dto.ProfileUrlResp, 0, len(profileUrls))
	for _, pUrl := range profileUrls {
		item := dto.ToProfileUrlDTO(&pUrl)
		res = append(res, item)
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) GetProfileUrlByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	profileUrl, err := h.usecase.GetProfileUrlByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	res := dto.ToProfileUrlDTO(&profileUrl)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (h *ProfileUrlHandler) UpdateProfileUrl(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	var req dto.ProfileUrlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := domain.ProfileUrlUpdateInput{
		ProfileID: &req.ProfileID,
		Label:     &req.Label,
		Url:       &req.Url,
	}

	if err := h.usecase.UpdateProfileUrl(ctx, id, input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "Profile url updated",
	})
}

func (h *ProfileUrlHandler) DeleteProfileUrl(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("profile_url_id")

	if err := h.usecase.DeleteProfileUrl(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "Profile url deleted",
	})
}

