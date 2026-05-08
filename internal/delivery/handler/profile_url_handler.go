package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProfileUrlHandler struct {
	usecase usecase.ProfileUrlUsecase
}

func NewProfileUrlHandler(u usecase.ProfileUrlUsecase) *ProfileUrlHandler {
	return &ProfileUrlHandler{usecase: u}
}

func (h *ProfileUrlHandler) CreateProfileUrl(w http.ResponseWriter, r *http.Request) {
	var req dto.ProfileUrlReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUrlInput{
		ProfileID: req.ProfileID,
		Label:     req.Label,
		Url:       req.Url,
	}

	profileUrl, err := h.usecase.CreateProfileUrl(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToProfileUrlDTO(profileUrl)
	ResponseJSON(w, r, http.StatusCreated, "created", res)
}

func (h *ProfileUrlHandler) GetProfileURL(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	profileUrls, err := h.usecase.GetProfileUrl(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := make([]dto.ProfileUrlResp, 0, len(profileUrls))
	for _, pUrl := range profileUrls {
		item := dto.ToProfileUrlDTO(&pUrl)
		res = append(res, item)
	}

	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *ProfileUrlHandler) GetProfileUrlByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profile_url_id")

	profileUrl, err := h.usecase.GetProfileUrlByID(r.Context(), id)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToProfileUrlDTO(&profileUrl)
	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *ProfileUrlHandler) UpdateProfileUrl(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profile_url_id")

	var req dto.ProfileUrlReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUrlUpdateInput{
		ProfileID: &req.ProfileID,
		Label:     &req.Label,
		Url:       &req.Url,
	}

	if err := h.usecase.UpdateProfileUrl(r.Context(), id, input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "Profile url updated")
}

func (h *ProfileUrlHandler) DeleteProfileUrl(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profile_url_id")

	if err := h.usecase.DeleteProfileUrl(r.Context(), id); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "Profile url deleted")
}
