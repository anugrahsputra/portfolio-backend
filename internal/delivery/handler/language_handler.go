package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type LanguageHandler struct {
	usecase usecase.LanguageUsecase
}

func NewLanguageHandler(u usecase.LanguageUsecase) *LanguageHandler {
	return &LanguageHandler{usecase: u}
}

func (h *LanguageHandler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	var req dto.LanguageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToLanguageInput(&req)
	language, err := h.usecase.CreateLanguage(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	res := dto.ToLanguageDTO(&language)
	ResponseJSON(w, r, http.StatusCreated, "success", res)
}

func (h *LanguageHandler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	languages, err := h.usecase.GetLanguages(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusBadRequest, "bad request")
		return
	}

	res := make([]dto.LanguageResp, 0, len(languages))
	for _, l := range languages {
		item := dto.ToLanguageDTO(&l)
		res = append(res, item)
	}

	ResponseJSON(w, r, http.StatusOK, "success", res)
}

func (h *LanguageHandler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	languageID := chi.URLParam(r, "language_id")

	var req dto.LanguageUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToLanguageUpdateInput(&req)

	if err := h.usecase.UpdateLanguage(r.Context(), languageID, input); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}

func (h *LanguageHandler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	languageID := chi.URLParam(r, "language_id")

	if err := h.usecase.DeleteLanguage(r.Context(), languageID); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}
