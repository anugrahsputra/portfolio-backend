package repository

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
)

type resumeRepository struct {
	db *db.Queries
}

func NewResumeRepository(database *config.Database) domain.ResumeRepository {
	return &resumeRepository{db: db.New(database.Pool)}
}

func (r *resumeRepository) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	idStr, _ := uuid.Parse(id)
	resume, err := r.db.GetResume(ctx, idStr)
	if err != nil {
		return nil, err
	}

	result, err := mapper.ToResumeDomain(resume)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
