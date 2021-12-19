package DemoController

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"runtime"
)

type Fast struct {

}

func (f *Fast) CreateLog(ctx *fasthttp.RequestCtx) {
	txt := ctx.FormValue("txt")
	fmt.Println(11, txt)
	fmt.Printf("finish all tasks, go num: %d\n", runtime.NumGoroutine())
}
