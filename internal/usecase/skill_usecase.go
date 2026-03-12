package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type SkillUsecase interface {
	CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error)
	GetSkills(ctx context.Context, profileID string) (domain.Skill, error)
	UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error
	DeleteSkill(ctx context.Context, id string) error
}

type skillUsecase struct {
	repo domain.SkillRepository
}

func NewSkillUsecase(r domain.SkillRepository) SkillUsecase {
	return &skillUsecase{repo: r}
}

func (u *skillUsecase) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	result, err := u.repo.CreateSkill(ctx, s)
	if err != nil {
		return domain.Skill{}, err
	}

	return result, nil
}

func (u *skillUsecase) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	result, err := u.repo.GetSkills(ctx, profileID)
	if err != nil {
		return domain.Skill{}, err
	}

	return result, nil
}

func (u *skillUsecase) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	if err := u.repo.UpdateSkill(ctx, id, s); err != nil {
		return err
	}

	return nil
}

func (u *skillUsecase) DeleteSkill(ctx context.Context, id string) error {
	if err := u.repo.DeleteSkill(ctx, id); err != nil {
		return err
	}

	return nil
}
