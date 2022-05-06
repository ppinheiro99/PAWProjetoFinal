package routes

import (
	"github.com/PAWProjetoFinal/backend/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Echo the data sent on get
// @Description Echo the data sent though the get request.
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Subject
// @Router /echo [get]
// @Param name query string false "string valid" minlength(1) maxlength(10)
// @Failure 404 "Not found"
func EchoRepeat(c *gin.Context) {
	controllers.Echo(c)
}

// @Summary Devolve todas as disciplinas
// @Description Exibe a lista de todas as disciplinas de um professor
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Subject
// @Router /Subject [get]
// @Failure 404 "Not found"
func GetAllSubject(c *gin.Context) {
	controllers.GetAllSubjects(c)
}

// @Summary Devolve uma disciplina por Id
// @Description Exibe os detalhes de uma disciplina pelo ID
// @ID get-Subject-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Subject ID"
// @Success 200 {object} model.Subject
// @Router /Subject/{id} [get]
// @Failure 404 "Not found"
func GetSubjectById(c *gin.Context) {
	controllers.GetSubjectByID(c)
}

// @Summary Atualiza uma disciplina
// @Description Atualiza uma disciplina sobre a utilização da aplicação
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param Subject body model.Subject true "Udpdate Subject"
// @Param id path int true "Subject ID"
// @Router /Subject/{id} [put]
// @Success 200 {object} model.Subject
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func UpdateSubject(c *gin.Context) {
	controllers.UpdateSubject(c)
}

// @Summary Exclui uma disciplina pelo ID
// @Description Exclui uma avaliação realizada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Subject ID"
// @Router /Subject/{id} [delete]
// @Success 200 {object} model.Subject
// @Failure 404 "Not found"
func DeleteSubject(c *gin.Context) {
	controllers.DeleteSubject(c)
}

// @Summary Adiciona uma disciplina
// @Description Cria uma disciplina sobre a utilização da aplicação
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param Subject body model.Subject true "Add Subject"
// @Router /Subject [post]
// @Success 201 {object} model.Subject
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AddSubject(c *gin.Context) {
	controllers.AddSubject(c)
}
