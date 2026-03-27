package config

import (
	"errors"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	Dialer *gomail.Dialer
}

func NewMail(cfg *Config) (*Mail, error) {
	if cfg.GmailUser == "" || cfg.GmailPass == "" {
		return nil, errors.New("gmail credentials are required")
	}

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		cfg.GmailUser,
		cfg.GmailPass,
	)

	return &Mail{Dialer: d}, nil
}
