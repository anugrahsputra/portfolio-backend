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
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
	}

	profile, err := r.db.CreateProfile(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}
	result := mapper.ToProfileDomainFromDB(profile)

	return &result, nil
}

func (r *profileRepository) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid profile id: %w", err)
	}

	profile, err := r.db.GetProfile(ctx, idStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	result, err := mapper.ToProfileDomain(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to map profile: %w", err)
	}

	return &result, nil
}

func (r *profileRepository) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid profile id: %w", err)
	}

	// Fetch-and-merge strategy for partial updates
	existing, err := r.db.GetProfile(ctx, idStr)
	if err != nil {
		return fmt.Errorf("failed to fetch existing profile for update: %w", err)
	}

	param := db.UpdateProfileParams{
		ID:      idStr,
		Name:    existing.Name,
		Title:   existing.Title,
		About:   existing.About,
		Address: existing.Address,
		Email:   existing.Email,
		Phone:   existing.Phone,
	}

	if p.Name != nil {
		param.Name = *p.Name
	}
	if p.Title != nil {
		param.Title = *p.Title
	}
	if p.About != nil {
		param.About = *p.About
	}
	if p.Address != nil {
		param.Address = *p.Address
	}
	if p.Email != nil {
		param.Email = *p.Email
	}
	if p.Phone != nil {
		param.Phone = *p.Phone
	}

	if err := r.db.UpdateProfile(ctx, param); err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}

func (r *profileRepository) DeleteProfile(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid profile id: %w", err)
	}

	if err := r.db.DeleteProfile(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	return nil
}
