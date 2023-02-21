package api

import (
	"go-starter/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine) {
	app.Use(gin.Recovery(), gin.Logger())

	user := handler.NewUser()
	app.GET("/user/list", user.List)
}
