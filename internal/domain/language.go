package domain

import "context"

type Language struct {
	ProfileID   string
	Language    string
	Proficiency string
}

type LanguageInput struct {
	ProfileID   string
	Language    string
	Proficiency string
}

type LanguageUpdateInput struct {
	Language    string
	Proficiency string
}

type LanguageRepository interface {
	CreateLanguage(ctx context.Context, l LanguageInput) error
	GetLanguages(ctx context.Context) ([]LanguageInput, error)
	UpdateLanguage(ctx context.Context, id string, l LanguageUpdateInput) error
	DeleteLanguage(ctx context.Context, id string) error
}
