package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PAWProjetoFinal/backend/model"
	"github.com/PAWProjetoFinal/backend/services"

	"github.com/gin-gonic/gin"
)

func GetPresentationById(c *gin.Context) {
	var Presentation model.Presentations
	id := c.Param("id")
	fmt.Printf("id: %s\n", id)
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
		fmt.Printf("entrei aqui")
		return
	}

	//c.JSON(http.StatusOK, Presentation)

	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Presentation})
}

func AddPresentation(c *gin.Context) {
	var Presentation model.Presentations

	Presentation.Name = c.Request.FormValue("name")
	if Presentation.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Name is required!"})
		return
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

	//Presentation.Teacher = c.Keys["username"].(string)

	services.Db.Save(&Presentation)
	if Presentation.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Couldn't Add Presentation!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": Presentation.ID})
}
