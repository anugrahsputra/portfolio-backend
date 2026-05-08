package handler

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type ProjectHandler struct {
	usecase usecase.ProjectUsecase
}

func NewProjectHandler(u usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{usecase: u}
}

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	ctx := c.Context()

	var req dto.ProjectReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
	}

	input := dto.ToProjectInput(&req)
	project, err := h.usecase.CreateProject(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to create project url: %v", err),
		})
	}

	resp := dto.ToProjectDTO(&project)
	return c.Status(http.StatusCreated).JSON(dto.Response{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    resp,
	})
}

func (h *ProjectHandler) GetProjects(c fiber.Ctx) error {
	ctx := c.Context()
	profileID := c.Params("profile_id")

	projects, err := h.usecase.GetProjects(ctx, profileID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Failed to get projects: %v", err),
		})
	}

	resp := make([]dto.ProjectResp, 0, len(projects))
	for _, project := range projects {
		item := dto.ToProjectDTO(&project)
		resp = append(resp, item)
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (h *ProjectHandler) UpdateProject(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("project_id")

	var req dto.ProjectUpdateReq
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.NoDataResponse{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
	}

	input := dto.ToProjectUpdateInput(&req)
	project, err := h.usecase.UpdateProject(ctx, id, input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to update project: %v", err),
		})
	}

	resp := dto.ToProjectDTO(&project)
	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (h *ProjectHandler) DeleteProject(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("project_id")

	if err := h.usecase.DeleteProject(ctx, id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.NoDataResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to delete project: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.Response{
		Status:  http.StatusOK,
		Message: "Success",
	})
}

