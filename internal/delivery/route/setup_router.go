package route

import (
	"net/http"
	"os"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
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

func SetupRouter(db *config.Database, mail *config.Mail, cfg *config.Config) *chi.Mux {
	env := os.Getenv("ENV")
	var allowOrigins []string

	if env == "development" {
		allowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	} else {
		allowOrigins = []string{"https://www.downormal.dev", "https://downormal.dev", "http://localhost:3000"}
	}

	r := chi.NewRouter()

	// Global Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}).Handler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		handler.ResponseError(w, r, http.StatusNotFound, "Endpoint not found")
	})

	profile := wireProfileRoute(db)
	profileUrl := wireProfileUrlRoute(db)
	experience := wireExperienceRoute(db)
	education := wireEducationRoute(db)
	skill := wireSkillRoute(db)
	language := wireLanguageRoute(db)
	project := wireProjectRoute(db)
	contactForm := wireContactFormRoute(db, mail, cfg)
	resume := wireResumeRoute(db)

	apiKey := os.Getenv("API_KEY")

	// API Group
	r.Route("/api/v1", func(r chi.Router) {
		ProfileRoute(r, profile, apiKey)
		ProfileUrlRoute(r, profileUrl, apiKey)
		ExperienceRoute(r, experience, apiKey)
		EducationRoute(r, education, apiKey)
		SkillRoute(r, skill, apiKey)
		LanguageRoute(r, language, apiKey)
		ProjectRoute(r, project, apiKey)
		ContactFormRoute(r, contactForm, apiKey)
		ResumeRoute(r, resume)
	})

	return r
}
