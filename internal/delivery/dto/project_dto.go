package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ProjectResp struct {
	ID            string   `json:"id"`
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
	Period        string   `json:"period"`
	Location      string   `json:"location"`
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
	Period        string   `json:"period"`
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
	Period        *string   `json:"period"`
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
		Period:        pr.Period,
		Location:      pr.Location,
	}
}

func ToProjectInput(pr *ProjectReq) domain.ProjectInput {
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
		Period:        pr.Period,
		Location:      pr.Location,
	}
}

func ToProjectUpdateInput(pr *ProjectUpdateReq) domain.ProjectUpdateInput {
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
		Period:        pr.Period,
		Location:      pr.Location,
	}
}
