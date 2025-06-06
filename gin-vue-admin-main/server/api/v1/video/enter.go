package video

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	VideoUserApi
	VideoWorksApi
	VideoTopicApi
}

var (
	videoUserService  = service.ServiceGroupApp.VideoServiceGroup.VideoUserService
	videoWorksService = service.ServiceGroupApp.VideoServiceGroup.VideoWorksService
	videoTopicService = service.ServiceGroupApp.VideoServiceGroup.VideoTopicService
)
