package Routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/Application"
	"server/Controllers"
)

func (app RouterApp) VisitorsRoutes() {
	app.Gin.GET("/api", func(c *gin.Context) {
		req := Application.NewRequest(c)
		req.Response(http.StatusOK, "Api is working well!", []interface{}{})
	})

	app.Gin.GET("/create-user", Controllers.CreateUser)
}
