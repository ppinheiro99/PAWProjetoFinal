package routes

import (
	"github.com/PAWProjetoFinal/backend/controllers"

	"github.com/gin-gonic/gin"
)

func GetPresentationById(c *gin.Context) {
	controllers.GetPresentationById(c)
}

func AddPresentation(c *gin.Context) {
	controllers.AddPresentation(c)
}

func GetQuestionsByPresentationId(c *gin.Context) {
	controllers.GetQuestionsByPresentationId(c)
}
