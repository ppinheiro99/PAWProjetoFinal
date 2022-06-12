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
	"unicode"

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
		if (currentSlice[0] == "questions.number") == true {
			q.QuestionNumber = currentSlice[1]
			i++
		}
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

func SubmitAnswer(c *gin.Context) {

	var doneAnswer model.DoneAnswers
	var question model.Questions

	if err := c.ShouldBindJSON(&doneAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	services.Db.Find(&question, "ID = ?", doneAnswer.QuestionId)
	if question.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "A question was not found with ID provided!"})
		return
	} else {
		if question.CorrectAnswer == doneAnswer.Answer {
			doneAnswer.Was_Right = true
		} else {
			doneAnswer.Was_Right = false
		}
	}
	services.Db.Save(&doneAnswer)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Answer Submitted with success!", "Student:": doneAnswer.StudentUsername})
	return
}
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func GetPresentationAnswers(c *gin.Context) {
	var Presentation model.Presentations
	var PresentationQuestions []model.PresentationAndQuestion
	var DoneAnswers []model.DoneAnswers
	var usr model.User

	id := c.Param("id")

	services.Db.First(&Presentation, id)
	if Presentation.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Presentation not found!"})
		return
	}

	services.Db.Where("presentation_id = ?", Presentation.ID).Find(&PresentationQuestions)
	if len(PresentationQuestions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Questions for Presentation not found!"})
	}

	services.Db.Find(&usr, "username = ?", c.Keys["username"].(string))

	//Primeiro vemos se é um professor ou aluno

	// role := strings.Split(c.Keys["username"].(string), "@")

	if IsLetter(c.Keys["username"].(string)) == true {
		var i = 0
		for i < len(PresentationQuestions) {
			var DoneAnswer model.DoneAnswers
			services.Db.Find(&DoneAnswer, PresentationQuestions[i].QuestionID)
			//fmt.Printf("ID: ", PresentationQuestions[i].QuestionID)
			if DoneAnswer.ID != 0 {
				DoneAnswers = append(DoneAnswers, DoneAnswer)
			}
			i++
		}
		if len(DoneAnswers) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "No Answers were given to this presentation"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": DoneAnswers})
		}
	} else {
		//if a student , only show their answers
		var i = 0
		for i < len(PresentationQuestions) {
			var DoneAnswer model.DoneAnswers
			services.Db.Find(&DoneAnswer, PresentationQuestions[i].QuestionID)
			//fmt.Printf("ID: ", PresentationQuestions[i].QuestionID)
			if DoneAnswer.ID != 0 {
				if DoneAnswer.StudentUsername == usr.Username {
					DoneAnswers = append(DoneAnswers, DoneAnswer)
				}

			}
			i++
		}
		if len(DoneAnswers) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "No Answers were given to this presentation"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": DoneAnswers})
		}
	}

}

type SubjectGrades struct {
	SubjectName string `json:"subjectName"`
	Grades      []Grades
}
type Grades struct {
	StudentUsername  string `json:"studentUsername"`
	PresentationName string `json:"presentationName"`
	Classification   int    `json:"classification"`
}
type Username struct {
	Username string `json:"username"`
}
type TeacherGrades struct {
	SubjectGrades []SubjectGrades
}

func GetClassificationByPresentation(c *gin.Context) {
	var Presentation model.Presentations
	var SubjectPresentations []model.SubjectPresentations
	var PresentationQuestions model.PresentationAndQuestion
	var DoneAnswers []model.DoneAnswers
	var Username Username
	var Subject []model.Subject
	var PresentationAndQuestion []model.PresentationAndQuestion
	var Grade Grades

	if err := c.ShouldBindJSON(&Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}
	if IsLetter(Username.Username) {
		services.Db.Find(&Subject, "teacher = ?", Username.Username)
		//Vamos buscar cada disciplina do professor
		var j = 0
		var SubjectGrades SubjectGrades
		var TeacherGrades TeacherGrades
		for j < len(Subject) {

			var k = 0
			var Users []model.User
			services.Db.Find(&SubjectPresentations, "subject_id = ?", Subject[j].ID)
			SubjectGrades.SubjectName = Subject[j].Name
			for k < len(SubjectPresentations) {
				services.Db.Find(&PresentationAndQuestion, "presentation_id = ?", SubjectPresentations[k].PresentationID)
				var kk = 0
				for kk < len(PresentationAndQuestion) {

					if PresentationAndQuestion[kk].PresentationID == SubjectPresentations[k].PresentationID {
						services.Db.Find(&Users)
						var ind = 0
						for ind < len(Users) {
							services.Db.Find(&DoneAnswers, "student_username = ?", Users[ind].Username)
							//Temos que saber a apresentaçao de cada resposta
							i := 0
							Grade.Classification = 0
							Grade.PresentationName = ""
							for i < len(DoneAnswers) {
								///Vamos buscar a apresentação que tem aquela pergunta
								services.Db.First(&PresentationQuestions, DoneAnswers[i].QuestionId)
								/// vamos buscar o nome da apresentaçao
								services.Db.Find(&Presentation, "id = ?", PresentationAndQuestion[kk].QuestionID)
								///Se for diferente entao estamos a classificar uma apresentação diferente
								if Grade.PresentationName != Presentation.Name {
									Grade.PresentationName = Presentation.Name
									if DoneAnswers[i].Was_Right {
										Grade.Classification += 1
									}
								} else {
									if DoneAnswers[i].Was_Right {
										Grade.Classification += 1
									}
								}
								Grade.StudentUsername = Users[ind].Username
								SubjectGrades.Grades = append(SubjectGrades.Grades, Grade)

								i++
							}

							ind += 1
						}

					}
					kk += 1
				}
				k += 1
			}
			j += 1
			TeacherGrades.SubjectGrades = append(TeacherGrades.SubjectGrades, SubjectGrades)
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": TeacherGrades})
		return

	} else {
		// vamos buscara s respostas do aluno
		services.Db.Find(&DoneAnswers, "student_username = ?", Username.Username)
		if len(DoneAnswers) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "No Answers were given to this presentation"})
			return
		}
		//Temos que saber a apresentaçao de cada resposta
		i := 0
		var Grade Grades
		var GradeArray []Grades
		Grade.Classification = 0
		Grade.PresentationName = ""

		for i < len(DoneAnswers) {
			///Vamos buscar a apresentação que tem aquela pergunta
			services.Db.First(&PresentationQuestions, DoneAnswers[i].QuestionId)
			/// vamos buscar o nome da apresentaçao
			services.Db.Find(&Presentation, "id = ?", PresentationQuestions.PresentationID)
			///Se for diferente entao estamos a classificar uma apresentação diferente
			if Grade.PresentationName != Presentation.Name {
				Grade.PresentationName = Presentation.Name
				if DoneAnswers[i].Was_Right {
					Grade.Classification += 1
				}
			} else {
				if DoneAnswers[i].Was_Right {
					Grade.Classification += 1
				}
			}

			GradeArray = append(GradeArray, Grade)

			i++
		}
		fmt.Printf("Inserted grade", GradeArray)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": GradeArray})
		return
	}

}
