package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	usecase usecase.ProjectUsecase
}

func NewProjectHandler(u usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{usecase: u}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req dto.ProjectReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToProjectInput(&req)
	project, err := h.usecase.CreateProject(r.Context(), input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := dto.ToProjectDTO(&project)
	ResponseJSON(w, r, http.StatusCreated, "created", resp)
}

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile_id")

	projects, err := h.usecase.GetProjects(r.Context(), profileID)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := make([]dto.ProjectResp, 0, len(projects))
	for _, project := range projects {
		item := dto.ToProjectDTO(&project)
		resp = append(resp, item)
	}

	ResponseJSON(w, r, http.StatusOK, "success", resp)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "project_id")

	var req dto.ProjectUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToProjectUpdateInput(&req)
	project, err := h.usecase.UpdateProject(r.Context(), id, input)
	if err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := dto.ToProjectDTO(&project)
	ResponseJSON(w, r, http.StatusOK, "success", resp)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "project_id")

	if err := h.usecase.DeleteProject(r.Context(), id); err != nil {
		ResponseError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(w, r, http.StatusOK, "success")
}
