package mapper

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToProjectDomain(pr db.Project) domain.Project {
	var ed *time.Time
	if pr.EndDate.Valid {
		ed = &pr.EndDate.Time
	}

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
		StartDate:     pr.StartDate.Time,
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}
}
