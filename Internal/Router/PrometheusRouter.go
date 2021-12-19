package Router

import (
	"blog_service/Api/BlogService/v1/Controller/PrometheusController"
	"github.com/gin-gonic/gin"
)

func PrometheusRouter(router *gin.Engine) {
	ins := new(PrometheusController.Prometheus)

	user := router.Group("/v1/prometheus")

	user.POST("/metrics", ins.GetsInfo)
}
