package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoTopicRouter struct {}

// InitVideoTopicRouter 初始化 videoTopic表 路由信息
func (s *VideoTopicRouter) InitVideoTopicRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	videoTopicRouter := Router.Group("videoTopic").Use(middleware.OperationRecord())
	videoTopicRouterWithoutRecord := Router.Group("videoTopic")
	videoTopicRouterWithoutAuth := PublicRouter.Group("videoTopic")
	{
		videoTopicRouter.POST("createVideoTopic", videoTopicApi.CreateVideoTopic)   // 新建videoTopic表
		videoTopicRouter.DELETE("deleteVideoTopic", videoTopicApi.DeleteVideoTopic) // 删除videoTopic表
		videoTopicRouter.DELETE("deleteVideoTopicByIds", videoTopicApi.DeleteVideoTopicByIds) // 批量删除videoTopic表
		videoTopicRouter.PUT("updateVideoTopic", videoTopicApi.UpdateVideoTopic)    // 更新videoTopic表
	}
	{
		videoTopicRouterWithoutRecord.GET("findVideoTopic", videoTopicApi.FindVideoTopic)        // 根据ID获取videoTopic表
		videoTopicRouterWithoutRecord.GET("getVideoTopicList", videoTopicApi.GetVideoTopicList)  // 获取videoTopic表列表
	}
	{
	    videoTopicRouterWithoutAuth.GET("getVideoTopicPublic", videoTopicApi.GetVideoTopicPublic)  // videoTopic表开放接口
	}
}
