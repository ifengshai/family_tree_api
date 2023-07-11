package api

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"go-admin/app/other/router"
)

func init() {
	fmt.Println(pkg.Green("api.other.init"))
	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}
