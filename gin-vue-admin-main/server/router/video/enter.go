package video

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	VideoUserRouter
	VideoWorksRouter
}

var (
	videoUserApi  = api.ApiGroupApp.VideoApiGroup.VideoUserApi
	videoWorksApi = api.ApiGroupApp.VideoApiGroup.VideoWorksApi
)
