package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type resumeUsecase struct {
	repo domain.ResumeRepository
}

func NewResumeUsecase(repo domain.ResumeRepository) domain.ResumeUsecase {
	return &resumeUsecase{repo: repo}
}

func (u *resumeUsecase) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	return u.repo.GetResume(ctx, id)
}
