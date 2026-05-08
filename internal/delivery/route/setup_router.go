package route

import (
	"errors"
	"os"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func wireProfileRoute(db *config.Database) *handler.ProfileHandler {
	profileRepo := repository.NewProfileRepository(db)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)
	profileHandler := handler.NewProfileHandler(profileUsecase)
	return profileHandler
}

func wireProfileUrlRoute(db *config.Database) *handler.ProfileUrlHandler {
	profileUrlRepo := repository.NewProfileUrlRepository(db)
	profileUrlUsecase := usecase.NewProfileUrlUsecase(profileUrlRepo)
	profileUrlHandler := handler.NewProfileUrlHandler(profileUrlUsecase)
	return profileUrlHandler
}

func wireExperienceRoute(db *config.Database) *handler.ExperienceHandler {
	experienceRepo := repository.NewExperienceRepository(db)
	experienceUsecase := usecase.NewExperienceUsecase(experienceRepo)
	experienceHandler := handler.NewExperienceHandler(experienceUsecase)
	return experienceHandler
}

func wireEducationRoute(db *config.Database) *handler.EducationHandler {
	educationRepo := repository.NewEducationRepository(db)
	educationUsecase := usecase.NewEducationUsecase(educationRepo)
	educationHandler := handler.NewEducationHandler(educationUsecase)
	return educationHandler
}

func wireSkillRoute(db *config.Database) *handler.SkillHandler {
	skillRepo := repository.NewSkillRepository(db)
	skillUsecase := usecase.NewSkillUsecase(skillRepo)
	skillHandler := handler.NewSkillHandler(skillUsecase)
	return skillHandler
}

func wireLanguageRoute(db *config.Database) *handler.LanguageHandler {
	languageRepo := repository.NewLanguageRepository(db)
	languageUsecase := usecase.NewLanguageUsecase(languageRepo)
	languageHandler := handler.NewLanguageHandler(languageUsecase)
	return languageHandler
}

func wireProjectRoute(db *config.Database) *handler.ProjectHandler {
	projectRepo := repository.NewProjectRepository(db)
	projectUsecase := usecase.NewProjectUsecase(projectRepo)
	projectHandler := handler.NewProjectHandler(projectUsecase)
	return projectHandler
}

func wireContactFormRoute(db *config.Database, mail *config.Mail, cfg *config.Config) *handler.ContactFormHandler {
	profileRepo := repository.NewProfileRepository(db)
	contactFormRepo := repository.NewEmailContactRepository(mail, cfg, profileRepo)
	contactFormUsecase := usecase.NewEmailContactUsecase(contactFormRepo)
	contactFormHandler := handler.NewContactFormHandler(contactFormUsecase)
	return contactFormHandler
}

func wireResumeRoute(db *config.Database) *handler.ResumeHandler {
	resumeRepo := repository.NewResumeRepository(db)
	resumeUsecase := usecase.NewResumeUsecase(resumeRepo)
	resumeHandler := handler.NewResumeHandler(resumeUsecase)
	return resumeHandler

}

func SetupRouter(db *config.Database, mail *config.Mail, cfg *config.Config) *fiber.App {
	env := os.Getenv("ENV")
	var allowOrigins []string

	if env == "development" {
		allowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	} else {
		allowOrigins = []string{"https://www.downormal.dev", "https://downormal.dev", "http://localhost:3000"}
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(dto.NoDataResponse{
				Status:  code,
				Message: err.Error(),
			})
		},
	})

	// Global Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	profile := wireProfileRoute(db)
	profileUrl := wireProfileUrlRoute(db)
	experience := wireExperienceRoute(db)
	education := wireEducationRoute(db)
	skill := wireSkillRoute(db)
	language := wireLanguageRoute(db)
	project := wireProjectRoute(db)
	contactForm := wireContactFormRoute(db, mail, cfg)
	resume := wireResumeRoute(db)

	// API Group
	apiKey := os.Getenv("API_KEY")
	api := app.Group("/api/v1")

	ProfileRoute(api, profile, apiKey)
	ProfileUrlRoute(api, profileUrl, apiKey)
	ExperienceRoute(api, experience, apiKey)
	EducationRoute(api, education, apiKey)
	SkillRoute(api, skill, apiKey)
	LanguageRoute(api, language, apiKey)
	ProjectRoute(api, project, apiKey)
	ContactFormRoute(api, contactForm, apiKey)
	ResumeRoute(api, resume)

	return app
}

