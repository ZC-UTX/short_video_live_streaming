package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoWorksRouter struct {}

// InitVideoWorksRouter 初始化 作品管理 路由信息
func (s *VideoWorksRouter) InitVideoWorksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	videoWorksRouter := Router.Group("videoWorks").Use(middleware.OperationRecord())
	videoWorksRouterWithoutRecord := Router.Group("videoWorks")
	videoWorksRouterWithoutAuth := PublicRouter.Group("videoWorks")
	{
		videoWorksRouter.POST("createVideoWorks", videoWorksApi.CreateVideoWorks)   // 新建作品管理
		videoWorksRouter.DELETE("deleteVideoWorks", videoWorksApi.DeleteVideoWorks) // 删除作品管理
		videoWorksRouter.DELETE("deleteVideoWorksByIds", videoWorksApi.DeleteVideoWorksByIds) // 批量删除作品管理
		videoWorksRouter.PUT("updateVideoWorks", videoWorksApi.UpdateVideoWorks)    // 更新作品管理
	}
	{
		videoWorksRouterWithoutRecord.GET("findVideoWorks", videoWorksApi.FindVideoWorks)        // 根据ID获取作品管理
		videoWorksRouterWithoutRecord.GET("getVideoWorksList", videoWorksApi.GetVideoWorksList)  // 获取作品管理列表
	}
	{
	    videoWorksRouterWithoutAuth.GET("getVideoWorksPublic", videoWorksApi.GetVideoWorksPublic)  // 作品管理开放接口
	}
}
