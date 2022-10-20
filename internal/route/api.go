package route

import (
	"go-starter/internal/handler/web"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.Use(gin.Recovery(), gin.Logger())

	user := web.NewUser()
	app.GET("/user/list", user.List)
}
