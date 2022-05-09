package main

import (
	"github.com/PAWProjetoFinal/backend/model"
	"github.com/PAWProjetoFinal/backend/routes"
	"github.com/PAWProjetoFinal/backend/services"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Presentations{})
	services.Db.AutoMigrate(&model.Subject{})
	services.Db.AutoMigrate(&model.PresentationAndQuestion{})
	services.Db.AutoMigrate(&model.SubjectPresentations{})
	services.Db.AutoMigrate(&model.Questions{})
	services.Db.AutoMigrate(&model.StudentsAndPresentationDone{})
	services.Db.AutoMigrate(&model.PresentationDone{})
	services.Db.AutoMigrate(&model.User{})

}

func main() {

	services.FormatSwagger()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(services.GinMiddleware("*"))

	// NO AUTH
	router.GET("/api/v1/echo", routes.EchoRepeat)

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	subject := router.Group("/api/v1/subject")
	subject.Use(services.AuthorizationRequired())
	{
		subject.POST("/", routes.AddSubject)
		subject.GET("/", routes.GetAllSubject)
		subject.GET("/:id", routes.GetSubjectById)
		subject.PUT("/:id", routes.UpdateSubject)
		subject.DELETE("/:id", routes.DeleteSubject)
	}

	presentations := router.Group("/api/v1/presentations")
	presentations.Use(services.AuthorizationRequired())
	{
		presentations.POST("/", routes.AddPresentation)
		presentations.GET("/:id", routes.GetPresentationById)
		presentations.DELETE("/:id", routes.DeletePresentationById)
		presentations.PUT("/:id", routes.UpdatePresentationById)
		presentations.GET("/:id/questions", routes.GetQuestionsByPresentationId)

	}

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(), routes.RefreshToken)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
