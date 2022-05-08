package controllers

import (
	"bytes"
	"fmt"
	"github.com/PAWProjetoFinal/backend/model"
	"github.com/PAWProjetoFinal/backend/services"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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

		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Something went wrong" + err.Error()})

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
