package repository

import (
	"context"
	"errors"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
)

type profileRepository struct {
	db *db.Queries
}

func NewProfileRepository(database *config.Database) domain.ProfileRepository {
	return &profileRepository{
		db: db.New(database.Pool),
	}
}

func (r *profileRepository) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	params := db.CreateProfileParams{
		Name:    p.Name,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
	}

	profile, err := r.db.CreateProfile(ctx, params)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	result := mapper.ToProfileDomainFromDB(profile)

	return &result, nil
}

func (r *profileRepository) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	idStr, _ := uuid.Parse(id)
	profile, err := r.db.GetProfile(ctx, idStr)
	if err != nil {
		return nil, err
	}

	result, err := mapper.ToProfileDomain(profile)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *profileRepository) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {

	idStr, err := uuid.Parse(id)
	if err != nil {
		return errors.New(err.Error())
	}

	param := db.UpdateProfileParams{
		ID:      idStr,
		Name:    *p.Name,
		About:   *p.About,
		Address: *p.Address,
		Email:   *p.Email,
		Phone:   *p.Email,
	}

	if err := r.db.UpdateProfile(ctx, param); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (r *profileRepository) DeleteProfile(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := r.db.DeleteProfile(ctx, idStr); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
