package route

import (
	"os"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
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

func SetupRouter(db *config.Database) *gin.Engine {
	env := os.Getenv("ENV")

	if env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
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

	// API Group
	api := route.Group("/api/v1")
	{
		ProfileRoute(api, profile)
		ProfileUrlRoute(api, profileUrl)
		ExperienceRoute(api, experience)
		EducationRoute(api, education)
		SkillRoute(api, skill)
		LanguageRoute(api, language)
		ProjectRoute(api, project)
	}

	return route
}
