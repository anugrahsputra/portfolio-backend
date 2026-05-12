package repository

import (
	"context"
	"fmt"

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
	idStr, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid resume id: %w", err)
	}

	resume, err := r.db.GetResume(ctx, idStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get resume: %w", err)
	}

	result, err := mapper.ToResumeDomain(resume)
	if err != nil {
		return nil, fmt.Errorf("failed to map resume: %w", err)
	}

	return &result, nil
}
