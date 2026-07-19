package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type languageUsecase struct {
	repo domain.LanguageRepository
}

func NewLanguageUsecase(repo domain.LanguageRepository) domain.LanguageUsecase {
	return &languageUsecase{repo: repo}
}

func (u *languageUsecase) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	return u.repo.CreateLanguage(ctx, l)
}

func (u *languageUsecase) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	return u.repo.GetLanguages(ctx, profileID)
}

func (u *languageUsecase) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	return u.repo.UpdateLanguage(ctx, id, l)
}

func (u *languageUsecase) DeleteLanguage(ctx context.Context, id string) error {
	return u.repo.DeleteLanguage(ctx, id)
}
