package dashboard

import (
	"github.com/gin-gonic/gin"
)

var DashboardResponse struct {
	countClass int
	countTeacher int
	countVisit int
	mostVisitedClass int
}

func getDashboard(c *gin.Context) {

}