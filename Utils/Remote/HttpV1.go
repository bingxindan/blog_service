package Remote

import (
	"context"
	"github.com/bingxindan/bxd_go_lib/logger"
	"github.com/kirinlabs/HttpRequest"
	"net/http"
)

var (
	httpClient *http.Client

	VideoUrl = "http://www.knowlet.cn:9998/video"
)

type HttpV1 struct {
}

type HttpV1Request struct {
	Uri     string                 `json:"uri"`
	Params  map[string]interface{} `json:"params"`
	Headers map[string]string      `json:"headers"`
	Auth    bool                   `json:"auth"`
}

type ResponseData struct {
	BodyData   string
	HeaderData http.Header
}

func (h *HttpV1) GetV1(ctx context.Context, request HttpV1Request) (string, error) {
	var (
		tag = "jz_api.Remote.HttpV1.GetV1"
	)
	req := HttpRequest.
		NewRequest().
		Debug(false).
		SetHeaders(request.Headers).
		SetTimeout(3)

	res, err := req.Get(request.Uri, request.Params)
	body, err := res.Body()
	if err != nil {
		logger.Ex(ctx, tag, "Get, req: %+v, ret: %+v, err: %+v", request, res, err)
		return "", err
	}
	return string(body), nil
}
