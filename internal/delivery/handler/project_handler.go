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
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToProjectInput(&req)
	project, err := h.usecase.CreateProject(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	resp := dto.ToProjectDTO(&project)
	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    resp,
	})
}

func (h *ProjectHandler) GetProjects(c *gin.Context) {
	ctx := c.Request.Context()
	profileID := c.Param("profile_id")

	projects, err := h.usecase.GetProjects(ctx, profileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	resp := make([]dto.ProjectResp, 0, len(projects))
	for _, project := range projects {
		item := dto.ToProjectDTO(&project)
		resp = append(resp, item)
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("project_id")

	var req dto.ProjectUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	input := dto.ToProjectUpdateInput(&req)
	project, err := h.usecase.UpdateProject(ctx, id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	resp := dto.ToProjectDTO(&project)
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("project_id")

	if err := h.usecase.DeleteProject(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.NoDataResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

