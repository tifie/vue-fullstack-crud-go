package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func SetupRouter(router *gin.Engine, injector *do.Injector) {
	todoRouter(router, injector)
}
