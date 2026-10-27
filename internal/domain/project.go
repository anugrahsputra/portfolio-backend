package domain

import (
	"context"
	"time"

	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
)

type Project struct {
	ID            string
	ProfileID     string
	Title         string
	Description   []string
	TechStacks    []string
	LiveDemoUrl   string
	GithubRepoUrl string
	IsLive        bool
	IsNda         bool
	IsComingSoon  bool
	IsFeatured    bool
	ImageUrl      string
	Company       string
	StartDate     time.Time
	EndDate       *time.Time
	IsPresent     bool
	Location      string
}

type ProjectInput struct {
	ProfileID     string
	Title         string
	Description   []string
	TechStacks    []string
	LiveDemoUrl   string
	GithubRepoUrl string
	IsLive        bool
	IsNda         bool
	IsComingSoon  bool
	IsFeatured    bool
	ImageUrl      string
	Company       string
	StartDate     time.Time
	EndDate       *time.Time
	IsPresent     bool
	Location      string
}

type ProjectUpdateInput struct {
	Title         *string
	Description   *[]string
	TechStacks    *[]string
	LiveDemoUrl   *string
	GithubRepoUrl *string
	IsLive        *bool
	IsNda         *bool
	IsComingSoon  *bool
	IsFeatured    *bool
	ImageUrl      *string
	Company       *string
	StartDate     *time.Time
	EndDate       *time.Time
	IsPresent     *bool
	Location      *string
}

// Merge applies the non-nil fields of in on top of p and returns the result.
// ProfileID is not mergeable: a project cannot change owner.
func (p Project) Merge(in ProjectUpdateInput) Project {
	p.Title = ptr.Or(in.Title, p.Title)
	p.Description = ptr.Or(in.Description, p.Description)
	p.TechStacks = ptr.Or(in.TechStacks, p.TechStacks)
	p.LiveDemoUrl = ptr.Or(in.LiveDemoUrl, p.LiveDemoUrl)
	p.GithubRepoUrl = ptr.Or(in.GithubRepoUrl, p.GithubRepoUrl)
	p.IsLive = ptr.Or(in.IsLive, p.IsLive)
	p.IsNda = ptr.Or(in.IsNda, p.IsNda)
	p.IsComingSoon = ptr.Or(in.IsComingSoon, p.IsComingSoon)
	p.IsFeatured = ptr.Or(in.IsFeatured, p.IsFeatured)
	p.ImageUrl = ptr.Or(in.ImageUrl, p.ImageUrl)
	p.Company = ptr.Or(in.Company, p.Company)
	p.StartDate = ptr.Or(in.StartDate, p.StartDate)
	p.IsPresent = ptr.Or(in.IsPresent, p.IsPresent)
	p.Location = ptr.Or(in.Location, p.Location)
	if in.EndDate != nil {
		p.EndDate = in.EndDate
	}
	return p
}

type ProjectUsecase interface {
	CreateProject(ctx context.Context, pr ProjectInput) (Project, error)
	GetProjects(ctx context.Context, profileID string) ([]Project, error)
	UpdateProject(ctx context.Context, id string, pr ProjectUpdateInput) (Project, error)
	DeleteProject(ctx context.Context, id string) error
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, pr ProjectInput) (Project, error)
	GetProjects(ctx context.Context, profileID string) ([]Project, error)
	GetProjectByID(ctx context.Context, id string) (Project, error)
	UpdateProject(ctx context.Context, pr Project) (Project, error)
	DeleteProject(ctx context.Context, id string) error
}


