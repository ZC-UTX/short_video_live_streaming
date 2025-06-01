package video

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/video"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type VideoTopicApi struct {}



// CreateVideoTopic 创建videoTopic表
// @Tags VideoTopic
// @Summary 创建videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoTopic true "创建videoTopic表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /videoTopic/createVideoTopic [post]
func (videoTopicApi *VideoTopicApi) CreateVideoTopic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var videoTopic video.VideoTopic
	err := c.ShouldBindJSON(&videoTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoTopic.CreatedBy = utils.GetUserID(c)
	err = videoTopicService.CreateVideoTopic(ctx,&videoTopic)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVideoTopic 删除videoTopic表
// @Tags VideoTopic
// @Summary 删除videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoTopic true "删除videoTopic表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /videoTopic/deleteVideoTopic [delete]
func (videoTopicApi *VideoTopicApi) DeleteVideoTopic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := videoTopicService.DeleteVideoTopic(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideoTopicByIds 批量删除videoTopic表
// @Tags VideoTopic
// @Summary 批量删除videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /videoTopic/deleteVideoTopicByIds [delete]
func (videoTopicApi *VideoTopicApi) DeleteVideoTopicByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := videoTopicService.DeleteVideoTopicByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideoTopic 更新videoTopic表
// @Tags VideoTopic
// @Summary 更新videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoTopic true "更新videoTopic表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /videoTopic/updateVideoTopic [put]
func (videoTopicApi *VideoTopicApi) UpdateVideoTopic(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var videoTopic video.VideoTopic
	err := c.ShouldBindJSON(&videoTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoTopic.UpdatedBy = utils.GetUserID(c)
	err = videoTopicService.UpdateVideoTopic(ctx,videoTopic)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideoTopic 用id查询videoTopic表
// @Tags VideoTopic
// @Summary 用id查询videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询videoTopic表"
// @Success 200 {object} response.Response{data=video.VideoTopic,msg=string} "查询成功"
// @Router /videoTopic/findVideoTopic [get]
func (videoTopicApi *VideoTopicApi) FindVideoTopic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	revideoTopic, err := videoTopicService.GetVideoTopic(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revideoTopic, c)
}
// GetVideoTopicList 分页获取videoTopic表列表
// @Tags VideoTopic
// @Summary 分页获取videoTopic表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoTopicSearch true "分页获取videoTopic表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /videoTopic/getVideoTopicList [get]
func (videoTopicApi *VideoTopicApi) GetVideoTopicList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo videoReq.VideoTopicSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := videoTopicService.GetVideoTopicInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetVideoTopicPublic 不需要鉴权的videoTopic表接口
// @Tags VideoTopic
// @Summary 不需要鉴权的videoTopic表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoTopic/getVideoTopicPublic [get]
func (videoTopicApi *VideoTopicApi) GetVideoTopicPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    videoTopicService.GetVideoTopicPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的videoTopic表接口信息",
    }, "获取成功", c)
}
