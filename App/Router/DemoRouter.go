package Router

import (
	"blog_service/App/Controller/DemoController"
	"github.com/gin-gonic/gin"
)

func DemoRouter(router *gin.Engine) {
	ins := new(DemoController.Demo)

	demo := router.Group("/")

	demo.GET("v1/demo/index/get", ins.GetById)
}
