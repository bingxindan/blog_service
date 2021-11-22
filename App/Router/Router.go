package Router

import "github.com/gin-gonic/gin"

func RouterRegister(router *gin.Engine) {
	DemoRouter(router)
}
