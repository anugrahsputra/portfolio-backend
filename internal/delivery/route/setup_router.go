package route

import (
	"os"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	apiKey := os.Getenv("API_KEY")

	profileRepo := repository.NewProfileRepository(db)
	profileUrlRepo := repository.NewProfileUrlRepository(db)
	experienceRepo := repository.NewExperienceRepository(db)
	educationRepo := repository.NewEducationRepository(db)
	skillRepo := repository.NewSkillRepository(db)
	languageRepo := repository.NewLanguageRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	resumeRepo := repository.NewResumeRepository(db)
	contactFormRepo := repository.NewEmailContactRepository(mail, cfg, profileRepo)

	profile := handler.NewProfileHandler(profileRepo)
	profileUrl := handler.NewProfileUrlHandler(profileUrlRepo)
	experience := handler.NewExperienceHandler(experienceRepo)
	education := handler.NewEducationHandler(educationRepo)
	skill := handler.NewSkillHandler(skillRepo)
	language := handler.NewLanguageHandler(languageRepo)
	project := handler.NewProjectHandler(projectRepo)
	resume := handler.NewResumeHandler(resumeRepo)
	contactForm := handler.NewContactFormHandler(contactFormRepo)

	api := route.Group("/api/v1")
	{
		r := api.Group("/profile")
		r.GET("/:id", profile.GetProfile)
		p := r.Group("")
		p.Use(middleware.AuthMiddleware(apiKey))
		p.GET("/", profile.GetProfiles)
		p.POST("/", profile.CreateProfile)
		p.PUT("/:id", profile.UpdateProfile)
		p.DELETE("/:id", profile.DeleteProfile)

		pu := api.Group("/profile-url")
		pu.GET("/:profile_id", profileUrl.GetProfileURL)
		pu.GET("/url/:profile_url_id", profileUrl.GetProfileUrlByID)
		pu2 := pu.Group("")
		pu2.Use(middleware.AuthMiddleware(apiKey))
		pu2.POST("/", profileUrl.CreateProfileUrl)
		pu2.PUT("/:profile_url_id", profileUrl.UpdateProfileUrl)
		pu2.DELETE("/:profile_url_id", profileUrl.DeleteProfileUrl)

		ex := api.Group("/experience")
		ex.GET("/:profile_id", experience.GetExperiences)
		ex2 := ex.Group("")
		ex2.Use(middleware.AuthMiddleware(apiKey))
		ex2.POST("/", experience.CreateExperience)
		ex2.PUT("/:experience_id", experience.UpdateExperience)
		ex2.DELETE("/:experience_id", experience.DeleteExperience)

		ed := api.Group("/education")
		ed.GET("/:profile_id", education.GetEducation)
		ed2 := ed.Group("")
		ed2.Use(middleware.AuthMiddleware(apiKey))
		ed2.POST("/", education.CreateEducation)
		ed2.PUT("/:education_id", education.UpdateEducation)
		ed2.DELETE("/:education_id", education.DeleteEducation)

		sk := api.Group("/skill")
		sk.GET("/:profile_id", skill.GetSkills)
		sk2 := sk.Group("")
		sk2.Use(middleware.AuthMiddleware(apiKey))
		sk2.POST("/", skill.CreateSkill)
		sk2.PUT("/:skill_id", skill.UpdateSkill)
		sk2.DELETE("/:skill_id", skill.DeleteSkill)

		la := api.Group("/language")
		la.GET("/:profile_id", language.GetLanguages)
		la2 := la.Group("")
		la2.Use(middleware.AuthMiddleware(apiKey))
		la2.POST("/", language.CreateLanguage)
		la2.PUT("/:language_id", language.UpdateLanguage)
		la2.DELETE("/:language_id", language.DeleteLanguage)

		pr := api.Group("/project")
		pr.GET("/:profile_id", project.GetProjects)
		pr2 := pr.Group("")
		pr2.Use(middleware.AuthMiddleware(apiKey))
		pr2.POST("", project.CreateProject)
		pr2.PUT("/:project_id", project.UpdateProject)
		pr2.DELETE("/:project_id", project.DeleteProject)

		se := api.Group("/send-email")
		se.Use(middleware.AuthMiddleware(apiKey))
		se.POST("/", contactForm.SendMail)

		re := api.Group("/resume")
		re.GET("/:profile_id", resume.GetResume)
	}

	return route
}
