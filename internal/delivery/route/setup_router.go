package route

import (
	"os"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func SetupRouter(db *config.Database, mail *config.Mail, cfg *config.Config) *gin.Engine {
	env := os.Getenv("ENV")
	var allowOrigins []string
	var allowMethods []string

	if env == "development" {
		gin.SetMode(gin.DebugMode)
		allowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
		allowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	} else {
		gin.SetMode(gin.ReleaseMode)
		allowOrigins = []string{"https://www.downormal.dev", "https://downormal.dev", "http://localhost:3000"}
		allowMethods = []string{"GET"}
	}

	route := gin.New()

	// Global Middlewares
	route.Use(middleware.RecoveryMiddleware())
	route.Use(gin.Logger())
	route.Use(middleware.SecurityMiddleware())
	route.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     allowMethods,
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
	api := route.Group("/api/v1")
	{
		ProfileRoute(api, profile, apiKey)
		ProfileUrlRoute(api, profileUrl, apiKey)
		ExperienceRoute(api, experience, apiKey)
		EducationRoute(api, education, apiKey)
		SkillRoute(api, skill, apiKey)
		LanguageRoute(api, language, apiKey)
		ProjectRoute(api, project, apiKey)
		ContactFormRoute(api, contactForm, apiKey)
		ResumeRoute(api, resume)
	}

	return route
}
