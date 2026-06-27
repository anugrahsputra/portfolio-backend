package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/google/uuid"
)

type resumeRepository struct {
	db *db.Queries
}

func NewResumeRepository(database *config.Database) domain.ResumeRepository {
	return &resumeRepository{db: db.New(database.Pool)}
}

func parseSlice[T any](input any) ([]T, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	var result []T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *resumeRepository) toDomain(res db.GetResumeRow) (domain.Resume, error) {
	urls, err := parseSlice[domain.ProfileUrl](res.Urls)
	if err != nil {
		return domain.Resume{}, err
	}
	skills, err := parseSlice[domain.Skill](res.Skills)
	if err != nil {
		return domain.Resume{}, err
	}
	languages, err := parseSlice[domain.Language](res.Languages)
	if err != nil {
		return domain.Resume{}, err
	}
	experiences, err := parseSlice[domain.Experience](res.Experiences)
	if err != nil {
		return domain.Resume{}, err
	}
	educations, err := parseSlice[domain.Education](res.Educations)
	if err != nil {
		return domain.Resume{}, err
	}
	projects, err := parseSlice[domain.Project](res.Projects)
	if err != nil {
		return domain.Resume{}, err
	}

	return domain.Resume{
		ID:          res.ID.String(),
		Name:        res.Name,
		Title:       res.Title,
		About:       res.About,
		Address:     res.Address,
		Email:       res.Email,
		Phone:       res.Phone,
		Url:         urls,
		Skills:      skills,
		Languages:   languages,
		Experiences: experiences,
		Educations:  educations,
		Projects:    projects,
	}, nil
}

func (r *resumeRepository) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid resume id: %w", err)
	}

	resume, err := r.db.GetResume(ctx, idStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get resume: %w", err)
	}

	d, err := r.toDomain(resume)
	if err != nil {
		return nil, fmt.Errorf("failed to map resume: %w", err)
	}

	return &d, nil
}
