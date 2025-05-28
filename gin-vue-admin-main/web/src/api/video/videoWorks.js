import service from '@/utils/request'
// @Tags VideoWorks
// @Summary 创建作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoWorks true "创建作品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /videoWorks/createVideoWorks [post]
export const createVideoWorks = (data) => {
  return service({
    url: '/videoWorks/createVideoWorks',
    method: 'post',
    data
  })
}

// @Tags VideoWorks
// @Summary 删除作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoWorks true "删除作品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoWorks/deleteVideoWorks [delete]
export const deleteVideoWorks = (params) => {
  return service({
    url: '/videoWorks/deleteVideoWorks',
    method: 'delete',
    params
  })
}

// @Tags VideoWorks
// @Summary 批量删除作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除作品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoWorks/deleteVideoWorks [delete]
export const deleteVideoWorksByIds = (params) => {
  return service({
    url: '/videoWorks/deleteVideoWorksByIds',
    method: 'delete',
    params
  })
}

// @Tags VideoWorks
// @Summary 更新作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoWorks true "更新作品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoWorks/updateVideoWorks [put]
export const updateVideoWorks = (data) => {
  return service({
    url: '/videoWorks/updateVideoWorks',
    method: 'put',
    data
  })
}

// @Tags VideoWorks
// @Summary 用id查询作品管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.VideoWorks true "用id查询作品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoWorks/findVideoWorks [get]
export const findVideoWorks = (params) => {
  return service({
    url: '/videoWorks/findVideoWorks',
    method: 'get',
    params
  })
}

// @Tags VideoWorks
// @Summary 分页获取作品管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取作品管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoWorks/getVideoWorksList [get]
export const getVideoWorksList = (params) => {
  return service({
    url: '/videoWorks/getVideoWorksList',
    method: 'get',
    params
  })
}

// @Tags VideoWorks
// @Summary 不需要鉴权的作品管理接口
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoWorksSearch true "分页获取作品管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoWorks/getVideoWorksPublic [get]
export const getVideoWorksPublic = () => {
  return service({
    url: '/videoWorks/getVideoWorksPublic',
    method: 'get',
  })
}
