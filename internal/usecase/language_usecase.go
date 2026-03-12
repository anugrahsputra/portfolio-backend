package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type LanguageUsecase interface {
	CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error)
	GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error)
	UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error
	DeleteLanguage(ctx context.Context, id string) error
}

type languageUsecase struct {
	repo domain.LanguageRepository
}

func NewLanguageUsecase(r domain.LanguageRepository) LanguageUsecase {
	return &languageUsecase{repo: r}
}

func (u *languageUsecase) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {

	result, err := u.repo.CreateLanguage(ctx, l)
	if err != nil {
		return domain.Language{}, err
	}

	return result, nil
}

func (u *languageUsecase) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	result, err := u.repo.GetLanguages(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *languageUsecase) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {

	if err := u.repo.UpdateLanguage(ctx, id, l); err != nil {
		return err
	}

	return nil
}

func (u *languageUsecase) DeleteLanguage(ctx context.Context, id string) error {

	if err := u.repo.DeleteLanguage(ctx, id); err != nil {
		return err
	}

	return nil
}
