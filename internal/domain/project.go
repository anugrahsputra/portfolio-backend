package domain

import "context"

type Project struct {
	ID            string
	ProfileID     string
	Title         string
	Description   string
	TechStacks    []string
	LiveDemoUrl   string
	GithubRepoUrl string
	IsLive        bool
	IsNda         bool
	IsFeatured    bool
	ImageUrl      string
	Company       string
	Period        string
	Location      string
}

type ProjectInput struct {
	ProfileID     string
	Title         string
	Description   string
	TechStacks    []string
	LiveDemoUrl   string
	GithubRepoUrl string
	IsLive        bool
	IsNda         bool
	IsFeatured    bool
	ImageUrl      string
	Company       string
	Period        string
	Location      string
}

type ProjectUpdateInput struct {
	Title         *string
	Description   *string
	TechStacks    *[]string
	LiveDemoUrl   *string
	GithubRepoUrl *string
	IsLive        *bool
	IsNda         *bool
	IsFeatured    *bool
	ImageUrl      *string
	Company       *string
	Period        *string
	Location      *string
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, pr ProjectInput) (Project, error)
	GetProjects(ctx context.Context, profileID string) ([]Project, error)
	UpdateProject(ctx context.Context, id string, pr ProjectUpdateInput) (Project, error)
	DeleteProject(ctx context.Context, id string) error
}
