package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "hello world",
	})
}
