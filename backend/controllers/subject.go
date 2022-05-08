package controllers

import (
	"net/http"
	"strconv"

	"github.com/PAWProjetoFinal/backend/model"
	"github.com/PAWProjetoFinal/backend/services"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	echo := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"echo": echo,
	})
}

func GetAllSubjects(c *gin.Context) {
	var Subjects []model.Subject

	services.Db.Find(&Subjects, "teacher = ?", c.Keys["username"].(string))

	if len(Subjects) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Subjects})
}

func GetSubjectByID(c *gin.Context) {
	var Subject model.Subject
	id := c.Param("id")

	services.Db.First(&Subject, id)
	if Subject.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Subject not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Subject})
}

func UpdateSubject(c *gin.Context) {
	var Subject model.Subject

	id := c.Param("id")
	services.Db.First(&Subject, id)

	if Subject.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Subject not found!"})
		return
	}

	if err := c.ShouldBindJSON(&Subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	services.Db.Save(Subject)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Update succeeded!"})
}

func AddSubject(c *gin.Context) {
	var Subject model.Subject

	if err := c.ShouldBindJSON(&Subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	Subject.Teacher = c.Keys["username"].(string)

	services.Db.Save(&Subject)
	if Subject.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Subject Already Exists!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": Subject.ID})
}

func DeleteSubject(c *gin.Context) {
	var Subject model.Subject
	var SubjectPresentations []model.SubjectPresentations
	var PresentationAndQuestions []model.PresentationAndQuestion

	Subject.Teacher = c.Keys["username"].(string)
	subjectid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check ID!"})
	}
	Subject.ID = uint(subjectid)

	services.Db.First(&Subject)

	if Subject.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Subject Not Found!"})
		return
	}

	///Before delete the subject, we need to delete the presentations and the questions
	services.Db.Find(&SubjectPresentations, "subject_id = ?", Subject.ID)
	for _, presentation := range SubjectPresentations {
		services.Db.Find(&PresentationAndQuestions, "presentation_id = ?", presentation.PresentationID)
		for _, question := range PresentationAndQuestions {
			services.Db.Delete(&question)
			services.Db.Exec("delete from questions where id = ? ", question.ID)
		}
		services.Db.Exec("delete from presentations where id = ? ", presentation.ID)
		services.Db.Exec("delete from presentation_and_questions where presentation_id = ? ", presentation.ID)
		services.Db.Exec("delete from subject_presentations where subject_id = ? ", Subject.ID)
	}

	services.Db.Exec("delete from subjects where id = ? and teacher = ?", Subject.ID, Subject.Teacher)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}
