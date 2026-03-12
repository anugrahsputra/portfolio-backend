package route

import (
	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
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

func SetupRouter(db *config.Database) *gin.Engine {
	route := gin.Default()

	profile := wireProfileRoute(db)
	profileUrl := wireProfileUrlRoute(db)
	experience := wireExperienceRoute(db)
	education := wireEducationRoute(db)
	skill := wireSkillRoute(db)
	language := wireLanguageRoute(db)

	// API Group
	api := route.Group("/api/v1")
	{
		ProfileRoute(api, profile)
		ProfileUrlRoute(api, profileUrl)
		ExperienceRoute(api, experience)
		EducationRoute(api, education)
		SkillRoute(api, skill)
		LanguageRoute(api, language)
	}

	return route
}
