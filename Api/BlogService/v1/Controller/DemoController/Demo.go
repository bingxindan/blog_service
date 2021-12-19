package DemoController

import (
	"blog_service/Internal/Service/DemoService"
	"blog_service/Internal/Struct/DemoStruct"
	"github.com/bingxindan/bxd_go_lib/api"
	"github.com/bingxindan/bxd_go_lib/gokit/network"
	"github.com/bingxindan/bxd_go_lib/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Demo struct {
	api.BaseController
}

// 消息回复
func (ctl *Demo) GetById(ctx *gin.Context) {
	var (
		tag     = ""
		goCtx   = network.TransferToContext(ctx)
		request = DemoStruct.IdxRequest{}
		demoSrv = DemoService.NewDemoService()
	)

	if err := ctx.ShouldBind(&request); err != nil {
		logger.Ex(goCtx, tag, "ShouldBind, req: %+v, err: %+v", request, err)
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	// 处理自动回复逻辑
	response, err := demoSrv.GetIndex(goCtx, request)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
