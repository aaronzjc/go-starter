package route

import (
	"go-starter/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.Use(gin.Recovery(), gin.Logger())

	test := new(handler.Test)
	app.GET("/hello", test.Get)

	user := new(handler.User)
	app.GET("/user/list", user.List)
}
