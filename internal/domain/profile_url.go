package domain

import "context"

type ProfileUrl struct {
	ID        string
	ProfileID string
	Label     string
	Url       string
}

type ProfileUrlInput struct {
	ProfileID string
	Label     string
	Url       string
}

type ProfileUrlUpdateInput struct {
	ProfileID *string
	Label     *string
	Url       *string
}

type ProfileUrRepository interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrl(ctx context.Context, id string) (ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, id string, pu ProfileUrlUpdateInput) error
	DeleteProfileUrl(ctx context.Context, id string) error
}
