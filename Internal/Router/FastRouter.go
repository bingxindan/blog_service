package Router

import (
	"blog_service/Api/BlogService/v1/Controller/DemoController"
	"github.com/buaazp/fasthttprouter"
)

func RegisterFast() *fasthttprouter.Router {
	router := fasthttprouter.New()

	srv := new(DemoController.Fast)
	router.POST("/write/log/create", srv.CreateLog)

	return router
}
