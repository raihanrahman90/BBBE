package article

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.Where("id = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return
	}

	if err := config.DB.Delete(&article).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return;
	}

	responseData := response(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	
}