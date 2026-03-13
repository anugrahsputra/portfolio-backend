package domain

import "context"

type Profile struct {
	ID      string
	Name    string
	About   string
	Address string
	Email   string
	Phone   string
	Url     []ProfileUrl `json:"urls"`
}

type ProfileInput struct {
	Name    string
	About   string
	Address string
	Email   string
	Phone   string
}

type ProfileUpdateInput struct {
	Name    *string
	About   *string
	Address *string
	Email   *string
	Phone   *string
}

type ProfileRepository interface {
	CreateProfile(ctx context.Context, p ProfileInput) (*Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	UpdateProfile(ctx context.Context, id string, p ProfileUpdateInput) error
	DeleteProfile(ctx context.Context, id string) error
}
