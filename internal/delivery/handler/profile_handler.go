package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(u usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecase: u}
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req dto.ProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileInput{
		Name:    req.Name,
		Title:   req.Title,
		About:   req.About,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	}

	profile, err := h.usecase.CreateProfile(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "Failed to create profile")
		return
	}

	res := dto.ToProfileDTO(profile)
	ResponseJSON(w, r, http.StatusCreated, "Success create profile", res)
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	profile, err := h.usecase.GetProfile(r.Context(), id)
	if err != nil {
		ResponseError(w, r, http.StatusNotFound, "Profile not found")
		return
	}

	res := dto.ToProfilePublicDTO(profile)
	ResponseJSON(w, r, http.StatusOK, "success get profile", res)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req dto.ProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := domain.ProfileUpdateInput{
		Name:    &req.Name,
		About:   &req.About,
		Address: &req.Address,
		Email:   &req.Email,
		Phone:   &req.Phone,
	}

	if err := h.usecase.UpdateProfile(r.Context(), id, input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	ResponseError(w, r, http.StatusOK, "success update profile")
}

func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.usecase.DeleteProfile(r.Context(), id); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "Failed to delete profile")
		return
	}

	ResponseError(w, r, http.StatusOK, "success delete profile")
}


