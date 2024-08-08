package controllers

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct{}

func (c *CourseController) GET(context *gin.Context) {
	id := context.Param("id")
	if id != "" {
		var course models.Course
		if err := inits.DATABASE.First(&course, id).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Error": "Course Not Found"})
			return
		}
		context.JSON(http.StatusOK, course)
	} else {
		var courses []models.Course
		if result := inits.DATABASE.Find(&courses); result.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}
		context.JSON(http.StatusOK, courses)
	}
}

func (c *CourseController) POST(context *gin.Context) {
	var reqBody struct {
		ID           uint
		CurriculumID string
		Course_No    string
		Course_Desc  string
		Lecture_Unit uint
		Semester     uint
		Year_Level   uint
	}

	context.Bind(&reqBody)
	course := models.Course{
		CurriculumID: reqBody.CurriculumID,
		Course_No:    reqBody.Course_No,
		Course_Desc:  reqBody.Course_Desc,
		Lecture_Unit: reqBody.Lecture_Unit,
		Semester:     reqBody.Semester,
		Year_Level:   reqBody.Year_Level,
	}
	if result := inits.DATABASE.Where(models.Course{Course_No: reqBody.Course_No, Course_Desc: reqBody.Course_Desc}).FirstOrCreate(&course); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new course record"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"course": course})

}

func (c *CourseController) DELETE(context *gin.Context) {
	id := context.Param("id")
	if err := inits.DATABASE.Delete(&models.Course{}, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Course Not Found"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
