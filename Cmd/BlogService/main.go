package main

import (
	"blog_service/Utils/Servers"
	"github.com/bingxindan/bxd_go_lib/tools/flagutil"
	"strconv"
)

// 服务应用管理周期，支持3种方式：
// 1、http
// 2、rpc
// 3、fasthttp
// 4、pprof监听端口，只允许内网访问
func main() {
	usr1 := *flagutil.GetUsr1()
	usr1Int, _ := strconv.Atoi(usr1)

	app, err := Servers.InitApp(usr1Int)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
