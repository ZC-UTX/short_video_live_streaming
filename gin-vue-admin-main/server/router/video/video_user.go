package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoUserRouter struct {}

// InitVideoUserRouter 初始化 用户管理 路由信息
func (s *VideoUserRouter) InitVideoUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	videoUserRouter := Router.Group("videoUser").Use(middleware.OperationRecord())
	videoUserRouterWithoutRecord := Router.Group("videoUser")
	videoUserRouterWithoutAuth := PublicRouter.Group("videoUser")
	{
		videoUserRouter.POST("createVideoUser", videoUserApi.CreateVideoUser)   // 新建用户管理
		videoUserRouter.DELETE("deleteVideoUser", videoUserApi.DeleteVideoUser) // 删除用户管理
		videoUserRouter.DELETE("deleteVideoUserByIds", videoUserApi.DeleteVideoUserByIds) // 批量删除用户管理
		videoUserRouter.PUT("updateVideoUser", videoUserApi.UpdateVideoUser)    // 更新用户管理
	}
	{
		videoUserRouterWithoutRecord.GET("findVideoUser", videoUserApi.FindVideoUser)        // 根据ID获取用户管理
		videoUserRouterWithoutRecord.GET("getVideoUserList", videoUserApi.GetVideoUserList)  // 获取用户管理列表
	}
	{
	    videoUserRouterWithoutAuth.GET("getVideoUserPublic", videoUserApi.GetVideoUserPublic)  // 用户管理开放接口
	}
}
