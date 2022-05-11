package routes

import (
	"github.com/PAWProjetoFinal/backend/controllers"

	"github.com/gin-gonic/gin"
)

func GetAllPresentation(c *gin.Context) {
	controllers.GetAllPresentation(c)
}

func GetPresentationById(c *gin.Context) {
	controllers.GetPresentationById(c)
}

func AddPresentation(c *gin.Context) {
	controllers.AddPresentation(c)
}

func GetQuestionsByPresentationId(c *gin.Context) {
	controllers.GetQuestionsByPresentationId(c)
}

func DeletePresentationById(c *gin.Context) {
	controllers.DeletePresentationById(c)
}

func UpdatePresentationById(c *gin.Context) {
	controllers.UpdatePresentationById(c)
}
