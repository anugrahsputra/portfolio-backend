package domain

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
)

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

// Merge applies the non-nil fields of in on top of pu and returns the result.
// ProfileID is not mergeable: a profile url cannot change owner.
func (pu ProfileUrl) Merge(in ProfileUrlUpdateInput) ProfileUrl {
	pu.Label = ptr.Or(in.Label, pu.Label)
	pu.Url = ptr.Or(in.Url, pu.Url)
	return pu
}

type ProfileUrlUsecase interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrlByID(ctx context.Context, id string) (ProfileUrl, error)
	GetProfileUrl(ctx context.Context, profileID string) ([]ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, id string, pu ProfileUrlUpdateInput) error
	DeleteProfileUrl(ctx context.Context, id string) error
}

type ProfileUrlRepository interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrlByID(ctx context.Context, id string) (ProfileUrl, error)
	GetProfileUrl(ctx context.Context, profileID string) ([]ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, pu ProfileUrl) error
	DeleteProfileUrl(ctx context.Context, id string) error
}


