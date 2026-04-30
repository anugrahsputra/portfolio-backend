package domain

import "context"

type Language struct {
	ID          string
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
	CreateLanguage(ctx context.Context, l LanguageInput) (Language, error)
	GetLanguages(ctx context.Context, profileID string) ([]Language, error)
	UpdateLanguage(ctx context.Context, id string, l LanguageUpdateInput) error
	DeleteLanguage(ctx context.Context, id string) error
}

func (l *Language) SetProfileID(profileID string) {
	l.ProfileID = profileID
}
