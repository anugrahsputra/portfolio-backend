package domain

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
)

type Profile struct {
	ID      string
	Name    string
	Title   string
	About   string
	Address string
	Email   string
	Phone   string
	Url     []ProfileUrl `json:"urls"`
}

type ProfileInput struct {
	Name    string
	Title   string
	About   string
	Address string
	Email   string
	Phone   string
}

type ProfileUpdateInput struct {
	Name    *string
	Title   *string
	About   *string
	Address *string
	Email   *string
	Phone   *string
}

// Merge applies the non-nil fields of in on top of p and returns the result.
func (p Profile) Merge(in ProfileUpdateInput) Profile {
	p.Name = ptr.Or(in.Name, p.Name)
	p.Title = ptr.Or(in.Title, p.Title)
	p.About = ptr.Or(in.About, p.About)
	p.Address = ptr.Or(in.Address, p.Address)
	p.Email = ptr.Or(in.Email, p.Email)
	p.Phone = ptr.Or(in.Phone, p.Phone)
	return p
}

type ProfileUsecase interface {
	CreateProfile(ctx context.Context, p ProfileInput) (*Profile, error)
	GetProfiles(ctx context.Context) ([]Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	UpdateProfile(ctx context.Context, id string, p ProfileUpdateInput) error
	DeleteProfile(ctx context.Context, id string) error
}

type ProfileRepository interface {
	CreateProfile(ctx context.Context, p ProfileInput) (*Profile, error)
	GetProfiles(ctx context.Context) ([]Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	UpdateProfile(ctx context.Context, p Profile) error
	DeleteProfile(ctx context.Context, id string) error
}
