package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSyPublicRouter)
}

// 需认证的路由代码
func registerSyPublicRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysPublic{}
	r := v1.Group("/public")
	{
		r.POST("/upload", api.UploadFile)
	}
}
