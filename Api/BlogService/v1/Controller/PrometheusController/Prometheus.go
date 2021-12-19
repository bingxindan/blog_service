package PrometheusController

import (
	"blog_service/Internal/Constant"
	"blog_service/Internal/Service/PrometheusService"
	"blog_service/Internal/Struct/PrometheusStruct"
	"github.com/bingxindan/bxd_go_lib/api"
	"github.com/bingxindan/bxd_go_lib/gokit/network"
	"github.com/bingxindan/bxd_go_lib/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Prometheus struct {
	api.BaseController
}

func (ctl *Prometheus) GetsInfo(c *gin.Context) {
	var (
		tag     = "jz_api.collection.controller.GetsInfo"
		ctx     = network.TransferToContext(c)

		request = PrometheusStruct.GetsInfoRequest{}
		srv     = PrometheusService.NewPrometheusService()
	)

	if err := c.ShouldBind(&request); err != nil {
		logger.Ex(ctx, tag, "ShouldBind, req: %+v, err: %+v", request, err)
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	response, err := srv.GetByIds(ctx, request)
	if err != nil {
		logger.Ex(ctx, tag, "GetByIds, req: %+v, err: %+v", request, err)
		c.JSON(http.StatusOK, network.Raw(Constant.StatSuccess, Constant.ResultFail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, network.Success(gin.H{"list": response}))
}
