package Router

import "github.com/gin-gonic/gin"

func RouterRegister(router *gin.Engine) *gin.Engine {
	// 示例
	DemoRouter(router)

	// 用户
	PrometheusRouter(router)

	return router
}
