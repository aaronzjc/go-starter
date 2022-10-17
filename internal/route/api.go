package route

import (
	"go-starter/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.Use(gin.Recovery(), gin.Logger())

	user := new(handler.User)
	app.GET("/user/list", user.List)
}
