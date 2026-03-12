package domain

import "context"

type Skill struct {
	ID           string
	ProfileID    string
	Tools        []string
	Technologies []string
	HardSkills   []string
	SoftSkills   []string
}

type SkillInput struct {
	ProfileID    string
	Tools        []string
	Technologies []string
	HardSkills   []string
	SoftSkills   []string
}

type SkillUpdateInput struct {
	Tools        []string
	Technologies []string
	HardSkills   []string
	SoftSkills   []string
}

type SkillRepository interface {
	CreateSkill(ctx context.Context, s SkillInput) (Skill, error)
	GetSkills(ctx context.Context, profileID string) (Skill, error)
	UpdateSkill(ctx context.Context, id string, s SkillUpdateInput) error
	DeleteSkill(ctx context.Context, id string) error
}
