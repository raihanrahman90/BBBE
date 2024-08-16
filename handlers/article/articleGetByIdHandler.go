package article

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetArticleById(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := config.DB.Where("id = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := response(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	return;
}