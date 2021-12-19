package DemoService

import (
	"blog_service/Internal/Model/DemoModel"
	"blog_service/Internal/Struct/DemoStruct"
	"context"
	"github.com/bingxindan/bxd_go_lib/logger"
	"sync"
)

type DemoService struct {
	demoMode *DemoModel.ArticleDao
}

var (
	serviceDemo *DemoService
	demoOnce    sync.Once
)

func NewDemoService() *DemoService {
	demoOnce.Do(func() {
		serviceDemo = &DemoService{
			demoMode: DemoModel.NewArticleDao(),
		}
	})
	return serviceDemo
}

// 查询索引
func (this *DemoService) GetIndex(ctx context.Context, request DemoStruct.IdxRequest) (response DemoModel.Article, err error) {
	var (
		tag = ""
	)

	// 通过ID查询文档
	response, err = this.demoMode.GetById(request.Id)
	if err != nil {
		logger.Ex(ctx, tag, "GetById, req: %+v, ret: %+v, err: %+v", request, response, err)
		return response, err
	}

	return response, nil
}
