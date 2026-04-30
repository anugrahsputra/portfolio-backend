package domain

import "context"

type Resume struct {
	ID          string
	Name        string
	Title       string
	About       string
	Address     string
	Email       string
	Phone       string
	Url         []ProfileUrl `json:"urls"`
	Skills      []Skill      `json:"skills"`
	Languages   []Language   `json:"languages"`
	Experiences []Experience `json:"experiences"`
	Educations  []Education  `json:"educations"`
	Projects    []Project    `json:"projects"`
}

type ResumeRepository interface {
	GetResume(ctx context.Context, id string) (*Resume, error)
}
