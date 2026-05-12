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

type skillRepository struct {
	db *db.Queries
}

func NewSkillRepository(database *config.Database) domain.SkillRepository {
	return &skillRepository{db: db.New(database.Pool)}
}

func (r *skillRepository) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	profileIDStr, err := uuid.Parse(s.ProfileID)
	if err != nil {
		return domain.Skill{}, fmt.Errorf("invalid profile id: %w", err)
	}

	param := db.CreateSkillParams{
		ProfileID:    profileIDStr,
		Tools:        s.Tools,
		Technologies: s.Technologies,
		HardSkills:   s.HardSkills,
		SoftSkills:   s.SoftSkills,
	}

	skill, err := r.db.CreateSkill(ctx, param)
	if err != nil {
		return domain.Skill{}, fmt.Errorf("failed to create skill: %w", err)
	}

	result := mapper.ToSkillDomain(&skill)
	return result, nil
}

func (r *skillRepository) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return domain.Skill{}, fmt.Errorf("invalid profile id: %w", err)
	}

	skill, err := r.db.GetSkillsByProfile(ctx, profileIDStr)
	if err != nil {
		return domain.Skill{}, fmt.Errorf("failed to get skills: %w", err)
	}

	result := mapper.ToSkillDomain(&skill)
	return result, nil
}

func (r *skillRepository) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid skill id: %w", err)
	}

	param := db.UpdateSkillParams{
		ID:           idStr,
		Tools:        s.Tools,
		Technologies: s.Technologies,
		HardSkills:   s.HardSkills,
		SoftSkills:   s.SoftSkills,
	}

	if _, err := r.db.UpdateSkill(ctx, param); err != nil {
		return fmt.Errorf("failed to update skill: %w", err)
	}

	return nil
}

func (r *skillRepository) DeleteSkill(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid skill id: %w", err)
	}

	if err := r.db.DeleteSkill(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete skill: %w", err)
	}

	return nil
}
