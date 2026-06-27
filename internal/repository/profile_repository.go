package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
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
	result := r.toDomainFromDB(profile)

	return &result, nil
}

func (r *profileRepository) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	profiles, err := r.db.GetProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get profiles")
	}

	var result []domain.Profile
	for _, profile := range profiles {
		d, err := r.toDomainFromGetProfilesRow(profile)
		if err != nil {
			return nil, fmt.Errorf("failed to map profile: %w", err)
		}
		result = append(result, d)
	}
	return result, nil
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

	d, err := r.toDomain(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to map profile: %w", err)
	}

	return &d, nil
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

func (r *profileRepository) toDomain(p db.GetProfileRow) (domain.Profile, error) {
	var urls []domain.ProfileUrl
	if p.Urls != nil {
		data, err := json.Marshal(p.Urls)
		if err != nil {
			return domain.Profile{}, err
		}
		if err := json.Unmarshal(data, &urls); err != nil {
			return domain.Profile{}, err
		}
	}

	return domain.Profile{
		ID:      p.ID.String(),
		Name:    p.Name,
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
		Url:     urls,
	}, nil
}

func (r *profileRepository) toDomainFromGetProfilesRow(p db.GetProfilesRow) (domain.Profile, error) {
	var urls []domain.ProfileUrl
	if p.Urls != nil {
		data, err := json.Marshal(p.Urls)
		if err != nil {
			return domain.Profile{}, err
		}
		if err := json.Unmarshal(data, &urls); err != nil {
			return domain.Profile{}, err
		}
	}

	return domain.Profile{
		ID:      p.ID.String(),
		Name:    p.Name,
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
		Url:     urls,
	}, nil
}

func (r *profileRepository) toDomainFromDB(p db.Profile) domain.Profile {
	return domain.Profile{
		ID:      p.ID.String(),
		Name:    p.Name,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
	}
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
