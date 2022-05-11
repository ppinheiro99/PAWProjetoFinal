package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PAWProjetoFinal/backend/model"
	"github.com/PAWProjetoFinal/backend/services"

	"github.com/gin-gonic/gin"
)

func GetAllPresentation(c *gin.Context) {
	var subjectsPresentations []model.SubjectPresentations
	id := c.Param("id")

	services.Db.Find(&subjectsPresentations, "subject_id = ?", id)

	if len(subjectsPresentations) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
		return
	}

	var responsePresentation []model.Presentations

	for i := 0; i < len(subjectsPresentations); i++ {
		var presentations model.Presentations
		services.Db.First(&presentations, subjectsPresentations[i].PresentationID)
		responsePresentation = append(responsePresentation, presentations)
	}

	services.Db.Find(&subjectsPresentations, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responsePresentation})
}

func GetPresentationById(c *gin.Context) {
	var Presentation model.Presentations
	id := c.Param("id")

	services.Db.First(&Presentation, id)
	if Presentation.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Presentation not found!"})
		return
	}
	err := os.WriteFile(Presentation.Name+".pdf", Presentation.PdfFile, 0644)
	if err == nil {
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename="+Presentation.Name+".pdf")
		c.Header("Content-Disposition", "inline;filename="+Presentation.Name+".pdf")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Cache-Control", "no-cache")
		c.Header("Content-Length", fmt.Sprintf("%d", len(Presentation.PdfFile)))
		c.File(Presentation.Name + ".pdf")
		e := os.Remove(Presentation.Name + ".pdf")
		if e != nil {
			log.Fatal(e)
		}

		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong" + err.Error()})
}
func DeletePresentationById(c *gin.Context) {
	var Presentation model.Presentations
	var PresentationAndQuestions []model.PresentationAndQuestion

	id := c.Param("id")

	services.Db.First(&Presentation, id)
	if Presentation.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Presentation not found!"})
		return
	}

	services.Db.Where("presentation_id = ?", id).Find(&PresentationAndQuestions)
	var i = 0
	for i < len(PresentationAndQuestions) {
		services.Db.Exec("DELETE FROM questions WHERE id = ?", PresentationAndQuestions[i].QuestionID)
		i++
	}
	services.Db.Exec("DELETE FROM subject_presentations WHERE presentation_id = ?", Presentation.ID)
	services.Db.Exec("DELETE FROM presentation_and_questions WHERE presentation_id = ?", Presentation.ID)
	services.Db.Exec("DELETE FROM presentations WHERE id = ?", Presentation.ID)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
	return

}
func UpdatePresentationById(c *gin.Context) {
	var Presentation model.Presentations
	var Subject model.Subject

	id := c.Param("id")
	services.Db.First(&Presentation, id)
	Presentation.Name = c.Request.FormValue("name")
	if Presentation.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Name is required!"})
		return
	}
	if c.Request.FormValue("questions") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Questions is required!"})
		return
	}
	if c.Request.FormValue("subjectid") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Subject is required!"})
		return
	}

	stringSlice := strings.Split(c.Request.FormValue("questions"), "&")
	var i = 0
	var question []model.Questions
	var q model.Questions
	for i < len(stringSlice) {
		currentSlice := strings.Split(stringSlice[i], "=")
		if (currentSlice[0] == "questions.question") == true {
			q.Question = currentSlice[1]
			i++
		}
		if (currentSlice[0] == "questions.answer") == true {

			q.Answers = q.Answers + "," + currentSlice[1]
			i++
		}
		if (currentSlice[0] == "questions.correct_answer") == true {
			q.CorrectAnswer = currentSlice[1]
			question = append(question, q)
			q.Answers = ""
			i++
		}
	}

	file, header, err := c.Request.FormFile("pdf_file")
	defer file.Close()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	fmt.Printf("File Name: %s\n", header.Filename)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error copying file : %s", err.Error()))
		return
	}

	Presentation.PdfFile = buf.Bytes()

	services.Db.Save(&Presentation)
	if Presentation.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Couldn't Add Presentation!"})
		return
	}

	i = 0
	for i < len(question) {
		q = question[i]
		services.Db.Save(&question[i])
		if question[i].ID != 0 {
			PandQ := model.PresentationAndQuestion{
				PresentationID: Presentation.ID,
				QuestionID:     question[i].ID,
			}
			services.Db.Save(&PandQ)
			if PandQ.ID == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})
				return
			}
			i++
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})
			return
		}
	}

	subjectid, err := strconv.Atoi(c.Request.FormValue("subjectid"))
	services.Db.Find(&Subject, subjectid)
	if Subject.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "SubjectId not valid!"})
		return
	}
	if err == nil {
		services.Db.Save(&model.SubjectPresentations{SubjectID: uint(subjectid), PresentationID: Presentation.ID})
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": Presentation.ID})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})
}

func GetQuestionsByPresentationId(c *gin.Context) {
	var Presentation model.Presentations
	var PresentationQuestions []model.PresentationAndQuestion
	var Questions []model.Questions
	id := c.Param("id")

	services.Db.First(&Presentation, id)
	if Presentation.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Presentation not found!"})
		return
	}
	//ids das perguntas para esta apresentação
	services.Db.Where("presentation_id = ?", Presentation.ID).Find(&PresentationQuestions)
	if len(PresentationQuestions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Questions for Presentation not found!"})
	}
	var i = 0
	for i < len(PresentationQuestions) {
		var Question model.Questions
		services.Db.First(&Question, PresentationQuestions[i].QuestionID)
		Questions = append(Questions, Question)
		i++
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Questions})

}

func AddPresentation(c *gin.Context) {
	var Presentation model.Presentations
	var Subject model.Subject
	Presentation.Name = c.Request.FormValue("name")
	if Presentation.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Name is required!"})
		return
	}
	if c.Request.FormValue("questions") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Questions is required!"})
		return
	}
	if c.Request.FormValue("subjectid") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Subject is required!"})
		return
	}

	stringSlice := strings.Split(c.Request.FormValue("questions"), "&")
	var i = 0
	var question []model.Questions
	var q model.Questions
	for i < len(stringSlice) {
		currentSlice := strings.Split(stringSlice[i], "=")
		if (currentSlice[0] == "questions.question") == true {
			q.Question = currentSlice[1]
			i++
		}
		if (currentSlice[0] == "questions.answer") == true {

			q.Answers = q.Answers + "," + currentSlice[1]
			i++
		}
		if (currentSlice[0] == "questions.correct_answer") == true {
			q.CorrectAnswer = currentSlice[1]
			question = append(question, q)
			q.Answers = ""
			i++
		}
	}

	file, header, err := c.Request.FormFile("pdf_file")
	defer file.Close()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	fmt.Printf("File Name: %s\n", header.Filename)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error copying file : %s", err.Error()))
		return
	}

	Presentation.PdfFile = buf.Bytes()

	services.Db.Save(&Presentation)
	if Presentation.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Couldn't Add Presentation!"})
		return
	}

	i = 0
	for i < len(question) {
		q = question[i]
		services.Db.Save(&question[i])
		if question[i].ID != 0 {
			PandQ := model.PresentationAndQuestion{
				PresentationID: Presentation.ID,
				QuestionID:     question[i].ID,
			}
			services.Db.Save(&PandQ)
			if PandQ.ID == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})
				return
			}
			i++
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})
			return
		}
	}

	subjectid, err := strconv.Atoi(c.Request.FormValue("subjectid"))
	services.Db.Find(&Subject, subjectid)
	if Subject.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "SubjectId not valid!"})
		return
	}
	if err == nil {
		services.Db.Save(&model.SubjectPresentations{SubjectID: uint(subjectid), PresentationID: Presentation.ID})
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": Presentation.ID})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong!"})

}
