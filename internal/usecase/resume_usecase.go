package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ResumeUsecase interface {
	GetResume(ctx context.Context, id string) (*domain.Resume, error)
}

type resumeUsecase struct {
	repo domain.ResumeRepository
}

func NewResumeUsecase(r domain.ResumeRepository) ResumeUsecase {
	return &resumeUsecase{repo: r}
}

func (u *resumeUsecase) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	result, err := u.repo.GetResume(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
