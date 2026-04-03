package repository

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
)

type profileUrlRepository struct {
	db *db.Queries
}

func NewProfileUrlRepository(database *config.Database) domain.ProfileUrlRepository {
	return &profileUrlRepository{db: db.New(database.Pool)}
}

func (r *profileUrlRepository) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	profileId, err := uuid.Parse(pu.ProfileID)
	if err != nil {
		return nil, err
	}
	param := db.CreateProfileURLParams{
		ProfileID: profileId,
		Label:     pu.Label,
		Url:       pu.Url,
	}

	profileUrl, err := r.db.CreateProfileURL(ctx, param)
	if err != nil {
		return nil, err
	}

	result := mapper.ToProfileURLDomain(profileUrl)
	return &result, nil

}

func (r *profileUrlRepository) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	profileIdStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, err
	}

	profileUrls, err := r.db.ListProfileURLs(ctx, profileIdStr)
	if err != nil {
		return nil, err
	}

	result := make([]domain.ProfileUrl, 0, len(profileUrls))
	for _, profileUrl := range profileUrls {
		item := mapper.ToProfileURLDomain(profileUrl)
		result = append(result, item)
	}

	return result, nil
}

func (r *profileUrlRepository) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.ProfileUrl{}, err
	}

	profileUrl, err := r.db.GetProfileURL(ctx, idStr)
	if err != nil {
		return domain.ProfileUrl{}, err
	}

	result := mapper.ToProfileURLDomain(profileUrl)
	return result, nil
}

func (r *profileUrlRepository) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	param := db.UpdateProfileURLParams{
		ID:    idStr,
		Label: *pu.Label,
		Url:   *pu.Url,
	}

	if _, err := r.db.UpdateProfileURL(ctx, param); err != nil {
		return err
	}

	return nil
}

func (r *profileUrlRepository) DeleteProfileUrl(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := r.db.DeleteProfileURL(ctx, idStr); err != nil {
		return err
	}

	return nil
}
