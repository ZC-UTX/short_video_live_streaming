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

type VideoUserApi struct {}



// CreateVideoUser 创建用户管理
// @Tags VideoUser
// @Summary 创建用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoUser true "创建用户管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /videoUser/createVideoUser [post]
func (videoUserApi *VideoUserApi) CreateVideoUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var videoUser video.VideoUser
	err := c.ShouldBindJSON(&videoUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoUser.CreatedBy = utils.GetUserID(c)
	err = videoUserService.CreateVideoUser(ctx,&videoUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVideoUser 删除用户管理
// @Tags VideoUser
// @Summary 删除用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoUser true "删除用户管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /videoUser/deleteVideoUser [delete]
func (videoUserApi *VideoUserApi) DeleteVideoUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := videoUserService.DeleteVideoUser(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVideoUserByIds 批量删除用户管理
// @Tags VideoUser
// @Summary 批量删除用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /videoUser/deleteVideoUserByIds [delete]
func (videoUserApi *VideoUserApi) DeleteVideoUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := videoUserService.DeleteVideoUserByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVideoUser 更新用户管理
// @Tags VideoUser
// @Summary 更新用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body video.VideoUser true "更新用户管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /videoUser/updateVideoUser [put]
func (videoUserApi *VideoUserApi) UpdateVideoUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var videoUser video.VideoUser
	err := c.ShouldBindJSON(&videoUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoUser.UpdatedBy = utils.GetUserID(c)
	err = videoUserService.UpdateVideoUser(ctx,videoUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVideoUser 用id查询用户管理
// @Tags VideoUser
// @Summary 用id查询用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户管理"
// @Success 200 {object} response.Response{data=video.VideoUser,msg=string} "查询成功"
// @Router /videoUser/findVideoUser [get]
func (videoUserApi *VideoUserApi) FindVideoUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	revideoUser, err := videoUserService.GetVideoUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revideoUser, c)
}
// GetVideoUserList 分页获取用户管理列表
// @Tags VideoUser
// @Summary 分页获取用户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoUserSearch true "分页获取用户管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /videoUser/getVideoUserList [get]
func (videoUserApi *VideoUserApi) GetVideoUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo videoReq.VideoUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := videoUserService.GetVideoUserInfoList(ctx,pageInfo)
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

// GetVideoUserPublic 不需要鉴权的用户管理接口
// @Tags VideoUser
// @Summary 不需要鉴权的用户管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoUser/getVideoUserPublic [get]
func (videoUserApi *VideoUserApi) GetVideoUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    videoUserService.GetVideoUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户管理接口信息",
    }, "获取成功", c)
}
