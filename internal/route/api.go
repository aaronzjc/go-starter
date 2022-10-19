package route

import (
	"go-starter/internal/handler/web"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.Use(gin.Recovery(), gin.Logger())

	user := new(web.User)
	app.GET("/user/list", user.List)
}
