package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToProjectDomain(pr db.Project) domain.Project {
	return domain.Project{
		ID:            pr.ID.String(),
		ProfileID:     pr.ProfileID.String(),
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
