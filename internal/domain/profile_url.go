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

type ProfileUrlRepository interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrlByID(ctx context.Context, id string) (ProfileUrl, error)
	GetProfileUrl(ctx context.Context, profileID string) ([]ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, id string, pu ProfileUrlUpdateInput) error
	DeleteProfileUrl(ctx context.Context, id string) error
}

func (p *ProfileUrl) SetProfileID(profileID string) {
	p.ProfileID = profileID
}
