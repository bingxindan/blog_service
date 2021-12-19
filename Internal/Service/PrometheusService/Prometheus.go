package PrometheusService

import (
	"blog_service/Internal/Model/PrometheusModel"
	"blog_service/Internal/Struct/PrometheusStruct"
	"context"
	"github.com/bingxindan/bxd_go_lib/logger"
	"github.com/pkg/errors"
	"sync"
)

type PrometheusService struct {
	prometheusModel *PrometheusModel.PrometheusDao
}

var (
	servicePrometheus *PrometheusService
	prometheusOnce    sync.Once
)

func NewPrometheusService() *PrometheusService {
	prometheusOnce.Do(func() {
		servicePrometheus = &PrometheusService{
			prometheusModel: PrometheusModel.NewPrometheusDao(),
		}
	})
	return servicePrometheus
}

// 查询索引
func (t *PrometheusService) GetByIds(ctx context.Context, request PrometheusStruct.GetsInfoRequest) (response []PrometheusModel.Prometheus, err error) {
	var (
		tag = "jz_api.prometheus.service.GetByIds"
	)

	// 通过ID查询文档
	response, err = t.prometheusModel.GetByIds([]string{request.PrometheusId})
	if err != nil {
		logger.Ex(ctx, tag, "GetByIds, req: %+v, ret: %+v, err: %+v", request, response, err)
		return response, errors.Wrap(err, "查询结果为空！")
	}

	return response, nil
}
