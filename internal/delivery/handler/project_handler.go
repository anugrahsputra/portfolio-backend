package handler

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	usecase usecase.ProjectUsecase
}

func NewProjectHandler(u usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{usecase: u}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.ProjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToProjectInput(&req)
	project, err := h.usecase.CreateProject(ctx, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := dto.ToProjectDTO(&project)
	ResponseJSON(c, http.StatusCreated, "created", resp)
}

func (h *ProjectHandler) GetProjects(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	projects, err := h.usecase.GetProjects(ctx, profileID)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := make([]dto.ProjectResp, 0, len(projects))
	for _, project := range projects {
		item := dto.ToProjectDTO(&project)
		resp = append(resp, item)
	}

	ResponseJSON(c, http.StatusOK, "success", resp)
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("project_id")

	var req dto.ProjectUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input := dto.ToProjectUpdateInput(&req)
	project, err := h.usecase.UpdateProject(ctx, id, input)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := dto.ToProjectDTO(&project)
	ResponseJSON(c, http.StatusOK, "success", resp)
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("project_id")

	if err := h.usecase.DeleteProject(ctx, id); err != nil {
		ResponseError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	ResponseError(c, http.StatusOK, "success")
}

