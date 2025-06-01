package video

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	VideoUserRouter
	VideoWorksRouter
	VideoTopicRouter
}

var (
	videoUserApi  = api.ApiGroupApp.VideoApiGroup.VideoUserApi
	videoWorksApi = api.ApiGroupApp.VideoApiGroup.VideoWorksApi
	videoTopicApi = api.ApiGroupApp.VideoApiGroup.VideoTopicApi
)
