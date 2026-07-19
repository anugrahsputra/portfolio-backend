package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type educationUsecase struct {
	repo domain.EducationRepository
}

func NewEducationUsecase(repo domain.EducationRepository) domain.EducationUsecase {
	return &educationUsecase{repo: repo}
}

func (u *educationUsecase) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	return u.repo.CreateEducation(ctx, e)
}

func (u *educationUsecase) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	return u.repo.GetEducations(ctx, profileID)
}

func (u *educationUsecase) GetEducationByID(ctx context.Context, id string) (domain.Education, error) {
	return u.repo.GetEducationByID(ctx, id)
}

func (u *educationUsecase) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	current, err := u.repo.GetEducationByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load education for update: %w", err)
	}

	return u.repo.UpdateEducation(ctx, current.Merge(e))
}

func (u *educationUsecase) DeleteEducation(ctx context.Context, id string) error {
	return u.repo.DeleteEducation(ctx, id)
}
