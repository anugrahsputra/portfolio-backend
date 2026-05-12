package repository

import (
	"context"
	"fmt"
	"html"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"gopkg.in/gomail.v2"
)

type emailContactRepository struct {
	dialer      *gomail.Dialer
	profileRepo domain.ProfileRepository
	from        string
}

func NewEmailContactRepository(mail *config.Mail, cfg *config.Config, profileRepo domain.ProfileRepository) domain.EmailContactRepository {
	return &emailContactRepository{dialer: mail.Dialer, from: cfg.GmailUser, profileRepo: profileRepo}
}

func (r *emailContactRepository) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	profile, err := r.profileRepo.GetProfile(ctx, form.ProfileID)
	if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", fmt.Sprintf("Portfolio <%s>", r.from))
	m.SetHeader("To", profile.Email)
	m.SetHeader("Reply-To", form.Email)
	m.SetHeader("Subject", fmt.Sprintf("[Portfolio] %s - %s", form.Name, form.Subject))

	body := emailTemplate(form)
	m.SetBody("text/html", body)

	if err := r.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func emailTemplate(form domain.EmailContactFormInput) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>New Contact Message</title>
</head>
<body style="margin:0;padding:0;background-color:#f5f5f5;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color:#f5f5f5;padding:48px 20px;">
    <tr>
      <td align="center">
        <table width="580" cellpadding="0" cellspacing="0" style="max-width:580px;width:100%%;">

          <!-- Header -->
          <tr>
            <td>
              <table width="100%%" cellpadding="0" cellspacing="0">
                <tr>
                  <td style="background-color:#0a0a0a;border-radius:8px 8px 0 0;padding:28px 32px 24px;">
                    <table width="100%%" cellpadding="0" cellspacing="0">
                      <tr>
                        <td>
                          <p style="margin:0 0 6px;font-size:11px;letter-spacing:2px;text-transform:uppercase;color:#666666;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;">downormal.dev</p>
                          <p style="margin:0;font-size:22px;font-weight:600;color:#ffffff;letter-spacing:-0.5px;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;">New Contact Message</p>
                        </td>
                        <td align="right" style="vertical-align:middle;">
                          <div style="display:inline-block;background-color:#1f1f1f;border:1px solid #2a2a2a;border-radius:20px;padding:6px 14px;">
                            <p style="margin:0;font-size:11px;color:#888888;letter-spacing:1px;text-transform:uppercase;">Portfolio</p>
                          </div>
                        </td>
                      </tr>
                    </table>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="background-color:#ffffff;border-left:1px solid #e8e8e8;border-right:1px solid #e8e8e8;padding:32px;">

              <!-- From -->
              <table width="100%%" cellpadding="0" cellspacing="0">
                <tr>
                  <td style="padding:16px 0;border-bottom:1px solid #f0f0f0;">
                    <p style="margin:0 0 6px;font-size:10px;letter-spacing:2px;text-transform:uppercase;color:#aaaaaa;">From</p>
                    <p style="margin:0 0 2px;font-size:15px;font-weight:600;color:#0a0a0a;">%s</p>
                    <p style="margin:0;font-size:13px;color:#888888;">%s</p>
                  </td>
                </tr>
              </table>

              <!-- Subject -->
              <table width="100%%" cellpadding="0" cellspacing="0">
                <tr>
                  <td style="padding:16px 0;border-bottom:1px solid #f0f0f0;">
                    <p style="margin:0 0 6px;font-size:10px;letter-spacing:2px;text-transform:uppercase;color:#aaaaaa;">Subject</p>
                    <p style="margin:0;font-size:15px;color:#1a1a1a;">%s</p>
                  </td>
                </tr>
              </table>

              <!-- Message -->
              <table width="100%%" cellpadding="0" cellspacing="0">
                <tr>
                  <td style="padding:24px 0 8px;">
                    <p style="margin:0 0 14px;font-size:10px;letter-spacing:2px;text-transform:uppercase;color:#aaaaaa;">Message</p>
                    <p style="margin:0;font-size:14px;line-height:1.85;color:#333333;white-space:pre-wrap;">%s</p>
                  </td>
                </tr>
              </table>

            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="background-color:#fafafa;border:1px solid #e8e8e8;border-top:none;border-radius:0 0 8px 8px;padding:16px 32px;">
              <table width="100%%" cellpadding="0" cellspacing="0">
                <tr>
                  <td>
                    <p style="margin:0;font-size:11px;color:#bbbbbb;">Hit reply to respond directly &mdash; reply-to is set to the sender.</p>
                  </td>
                  <td align="right">
                    <p style="margin:0;font-size:11px;color:#cccccc;">downormal.dev</p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
</body>
</html>`,
		html.EscapeString(form.Name),
		html.EscapeString(form.Email),
		html.EscapeString(form.Subject),
		html.EscapeString(form.Message),
	)

}
