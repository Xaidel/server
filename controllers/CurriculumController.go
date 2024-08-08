package controllers

import (
	"csprobe/server/inits"
	"csprobe/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurriculumController struct{}

func (c *CurriculumController) GET(context *gin.Context) {
	id := context.Param("id")
	if id != "" {
		var curriculum models.Curriculum
		if err := inits.DATABASE.First(&curriculum, "id = ?", id).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Curriculum Not Found"})
			return
		}
		context.JSON(http.StatusOK, curriculum)
	} else {
		var curriculums []models.Curriculum
		if result := inits.DATABASE.Find(&curriculums); result.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		context.JSON(http.StatusOK, curriculums)
	}
}

func (c *CurriculumController) POST(context *gin.Context) {
	var reqBody struct {
		ID              string
		Program_Code    string
		Revision_Number uint
		Effectivity_Sem uint
		Effectivity_SY  uint
		CMO_Ref         string
		IsActive        bool
	}

	context.Bind(&reqBody)
	curriculum := models.Curriculum{
		ID:              reqBody.ID,
		Program_Code:    reqBody.Program_Code,
		Revision_Number: reqBody.Revision_Number,
		Effectivity_Sem: reqBody.Effectivity_Sem,
		Effectivity_SY:  reqBody.Effectivity_SY,
		CMO_Ref:         reqBody.CMO_Ref,
		IsActive:        reqBody.IsActive,
	}
	if result := inits.DATABASE.Where(models.Curriculum{ID: reqBody.ID}).FirstOrCreate(&curriculum); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new curriculum record"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": curriculum})
}

func (c *CurriculumController) DELETE(context *gin.Context) {
	id := context.Param("id")
	if err := inits.DATABASE.Delete(&models.Curriculum{}, "id = ?", id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Curriculum Not Found"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
