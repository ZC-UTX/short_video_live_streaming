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

type VideoWorksApi struct {}



// CreateVideoWorks 创建作品管理
// @Tags VideoWorks
// @Summary 创建作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoWorks true "创建作品管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /videoWorks/createVideoWorks [post]
func (videoWorksApi *VideoWorksApi) CreateVideoWorks(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var videoWorks video.VideoWorks
	err := c.ShouldBindJSON(&videoWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoWorks.CreatedBy = utils.GetUserID(c)
	err = videoWorksService.CreateVideoWorks(ctx,&videoWorks)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVideoWorks 删除作品管理
// @Tags VideoWorks
// @Summary 删除作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoWorks true "删除作品管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /videoWorks/deleteVideoWorks [delete]
func (videoWorksApi *VideoWorksApi) DeleteVideoWorks(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := videoWorksService.DeleteVideoWorks(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideoWorksByIds 批量删除作品管理
// @Tags VideoWorks
// @Summary 批量删除作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /videoWorks/deleteVideoWorksByIds [delete]
func (videoWorksApi *VideoWorksApi) DeleteVideoWorksByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := videoWorksService.DeleteVideoWorksByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideoWorks 更新作品管理
// @Tags VideoWorks
// @Summary 更新作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoWorks true "更新作品管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /videoWorks/updateVideoWorks [put]
func (videoWorksApi *VideoWorksApi) UpdateVideoWorks(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var videoWorks video.VideoWorks
	err := c.ShouldBindJSON(&videoWorks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoWorks.UpdatedBy = utils.GetUserID(c)
	err = videoWorksService.UpdateVideoWorks(ctx,videoWorks)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideoWorks 用id查询作品管理
// @Tags VideoWorks
// @Summary 用id查询作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询作品管理"
// @Success 200 {object} response.Response{data=video.VideoWorks,msg=string} "查询成功"
// @Router /videoWorks/findVideoWorks [get]
func (videoWorksApi *VideoWorksApi) FindVideoWorks(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	revideoWorks, err := videoWorksService.GetVideoWorks(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revideoWorks, c)
}
// GetVideoWorksList 分页获取作品管理列表
// @Tags VideoWorks
// @Summary 分页获取作品管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoWorksSearch true "分页获取作品管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /videoWorks/getVideoWorksList [get]
func (videoWorksApi *VideoWorksApi) GetVideoWorksList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo videoReq.VideoWorksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := videoWorksService.GetVideoWorksInfoList(ctx,pageInfo)
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

// GetVideoWorksPublic 不需要鉴权的作品管理接口
// @Tags VideoWorks
// @Summary 不需要鉴权的作品管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoWorks/getVideoWorksPublic [get]
func (videoWorksApi *VideoWorksApi) GetVideoWorksPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    videoWorksService.GetVideoWorksPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的作品管理接口信息",
    }, "获取成功", c)
}
