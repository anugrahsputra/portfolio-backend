package dto

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ProjectResp struct {
	ID            string    `json:"id"`
	ProfileID     string    `json:"profile_id"`
	Title         string    `json:"title"`
	Description   []string  `json:"description"`
	TechStacks    []string  `json:"tech_stacks"`
	LiveDemoUrl   string    `json:"live_demo_url"`
	GithubRepoUrl string    `json:"github_repo_url"`
	IsLive        bool      `json:"is_live"`
	IsNda         bool      `json:"is_nda"`
	IsFeatured    bool      `json:"is_featured"`
	ImageUrl      string    `json:"image_url"`
	Company       string    `json:"company"`
	StartDate     time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	IsPresent     bool      `json:"is_present"`
	Location      string    `json:"location"`
}

type ProjectReq struct {
	ProfileID     string   `json:"profile_id"`
	Title         string   `json:"title"`
	Description   []string `json:"description"`
	TechStacks    []string `json:"tech_stacks"`
	LiveDemoUrl   string   `json:"live_demo_url"`
	GithubRepoUrl string   `json:"github_repo_url"`
	IsLive        bool     `json:"is_live"`
	IsNda         bool     `json:"is_nda"`
	IsFeatured    bool     `json:"is_featured"`
	ImageUrl      string   `json:"image_url"`
	Company       string   `json:"company"`
	StartDate     string   `json:"start_date"`
	EndDate       string   `json:"end_date"`
	IsPresent     bool     `json:"is_present"`
	Location      string   `json:"location"`
}

type ProjectUpdateReq struct {
	Title         *string   `json:"title"`
	Description   *[]string `json:"description"`
	TechStacks    *[]string `json:"tech_stacks"`
	LiveDemoUrl   *string   `json:"live_demo_url"`
	GithubRepoUrl *string   `json:"github_repo_url"`
	IsLive        *bool     `json:"is_live"`
	IsNda         *bool     `json:"is_nda"`
	IsFeatured    *bool     `json:"is_featured"`
	ImageUrl      *string   `json:"image_url"`
	Company       *string   `json:"company"`
	StartDate     *string   `json:"start_date"`
	EndDate       *string   `json:"end_date"`
	IsPresent     *bool     `json:"is_present"`
	Location      *string   `json:"location"`
}

func ToProjectDTO(pr *domain.Project) ProjectResp {
	return ProjectResp{
		ID:            pr.ID,
		ProfileID:     pr.ProfileID,
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     pr.StartDate,
		EndDate:       pr.EndDate,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}
}

func ToProjectInput(pr *ProjectReq) domain.ProjectInput {
	sd, _ := time.Parse("2006-01-02", pr.StartDate)
	var ed *time.Time
	if pr.EndDate != "" {
		parsed, err := time.Parse("2006-01-02", pr.EndDate)
		if err == nil {
			ed = &parsed
		}
	}

	return domain.ProjectInput{
		ProfileID:     pr.ProfileID,
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     sd,
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}
}

func ToProjectUpdateInput(pr *ProjectUpdateReq) domain.ProjectUpdateInput {
	var sd *time.Time
	if pr.StartDate != nil {
		parsed, err := time.Parse("2006-01-02", *pr.StartDate)
		if err == nil {
			sd = &parsed
		}
	}

	var ed *time.Time
	if pr.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *pr.EndDate)
		if err == nil {
			ed = &parsed
		}
	}

	return domain.ProjectUpdateInput{
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     sd,
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}
}
