package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct{}

func (c *Test) Get(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
