package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/pkg/parser"
)

func ToResumeDomain(r db.GetResumeRow) (domain.Resume, error) {
	id := r.ID.String()

	urls, err := parser.JsonSliceParser[domain.ProfileUrl](r.Urls, id)
	if err != nil {
		return domain.Resume{}, err
	}

	skills, err := parser.JsonSliceParser[domain.Skill](r.Skills, id)
	if err != nil {
		return domain.Resume{}, err
	}

	languages, err := parser.JsonSliceParser[domain.Language](r.Languages, id)
	if err != nil {
		return domain.Resume{}, err
	}

	experiences, err := parser.JsonSliceParser[domain.Experience](r.Experiences, id)
	if err != nil {
		return domain.Resume{}, err
	}

	educations, err := parser.JsonSliceParser[domain.Education](r.Educations, id)
	if err != nil {
		return domain.Resume{}, err
	}

	projects, err := parser.JsonSliceParser[domain.Project](r.Projects, id)
	if err != nil {
		return domain.Resume{}, err
	}

	return domain.Resume{
		ID:          r.ID.String(),
		Name:        r.Name,
		Title:       r.Title,
		About:       r.About,
		Address:     r.Address,
		Email:       r.Email,
		Phone:       r.Phone,
		Url:         urls,
		Skills:      skills,
		Languages:   languages,
		Experiences: experiences,
		Educations:  educations,
		Projects:    projects,
	}, nil
}
