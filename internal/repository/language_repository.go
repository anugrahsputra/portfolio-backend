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

type languageRepository struct {
	db *db.Queries
}

func NewLanguageRepository(database *config.Database) domain.LanguageRepository {
	return &languageRepository{db: db.New(database.Pool)}
}

func (r *languageRepository) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	profileID, err := uuid.Parse(l.ProfileID)
	if err != nil {
		return domain.Language{}, fmt.Errorf("invalid profile id: %w", err)
	}

	param := db.CreateLanguageParams{
		ProfileID:   profileID,
		Language:    l.Language,
		Proficiency: db.ProficiencyLevel(l.Proficiency),
	}

	language, err := r.db.CreateLanguage(ctx, param)
	if err != nil {
		return domain.Language{}, fmt.Errorf("failed to create language: %w", err)
	}

	result := mapper.ToLanguageDomain(&language)
	return result, nil
}

func (r *languageRepository) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, fmt.Errorf("invalid profile id: %w", err)
	}

	languages, err := r.db.ListLanguages(ctx, profileIDStr)
	if err != nil {
		return nil, fmt.Errorf("failed to list languages: %w", err)
	}

	var result []domain.Language
	for _, language := range languages {
		item := mapper.ToLanguageDomain(&language)
		result = append(result, item)
	}

	return result, nil
}

func (r *languageRepository) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid language id: %w", err)
	}

	param := db.UpdateLanguageParams{
		ID:          idStr,
		Language:    l.Language,
		Proficiency: db.ProficiencyLevel(l.Proficiency),
	}

	if _, err := r.db.UpdateLanguage(ctx, param); err != nil {
		return fmt.Errorf("failed to update language: %w", err)
	}

	return nil
}

func (r *languageRepository) DeleteLanguage(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid language id: %w", err)
	}

	if err := r.db.DeleteLanguage(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete language: %w", err)
	}

	return nil
}
