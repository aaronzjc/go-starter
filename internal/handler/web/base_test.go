package web

import (
	"go-starter/internal/constant"
	"go-starter/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResp(t *testing.T) {
	assert := assert.New(t)
	h := func(c *gin.Context) {
		Resp(c, 0, "ok", nil)
	}

	resp := test.NewRequest(t).Handler(h).Get("/").Exec()
	assert.Equal(200, resp.Code())
	errno, errmsg, _, _ := resp.TryDecode()
	assert.Equal(errno, constant.ERR_OK)
	assert.Equal(errmsg, "ok")
}
