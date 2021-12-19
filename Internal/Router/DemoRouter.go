package Router

import (
	"blog_service/Api/BlogService/v1/Controller/DemoController"
	"github.com/gin-gonic/gin"
)

func DemoRouter(router *gin.Engine) {
	ins := new(DemoController.Demo)

	demo := router.Group("/")

	demo.GET("v1/demo/index/get", ins.GetById)
}
