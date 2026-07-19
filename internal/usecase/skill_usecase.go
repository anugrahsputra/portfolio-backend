package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type skillUsecase struct {
	repo domain.SkillRepository
}

func NewSkillUsecase(repo domain.SkillRepository) domain.SkillUsecase {
	return &skillUsecase{repo: repo}
}

func (u *skillUsecase) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	return u.repo.CreateSkill(ctx, s)
}

func (u *skillUsecase) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	return u.repo.GetSkills(ctx, profileID)
}

func (u *skillUsecase) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	return u.repo.UpdateSkill(ctx, id, s)
}

func (u *skillUsecase) DeleteSkill(ctx context.Context, id string) error {
	return u.repo.DeleteSkill(ctx, id)
}
