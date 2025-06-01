import service from '@/utils/request'
// @Tags VideoTopic
// @Summary 创建videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoTopic true "创建videoTopic表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /videoTopic/createVideoTopic [post]
export const createVideoTopic = (data) => {
  return service({
    url: '/videoTopic/createVideoTopic',
    method: 'post',
    data
  })
}

// @Tags VideoTopic
// @Summary 删除videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoTopic true "删除videoTopic表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoTopic/deleteVideoTopic [delete]
export const deleteVideoTopic = (params) => {
  return service({
    url: '/videoTopic/deleteVideoTopic',
    method: 'delete',
    params
  })
}

// @Tags VideoTopic
// @Summary 批量删除videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除videoTopic表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoTopic/deleteVideoTopic [delete]
export const deleteVideoTopicByIds = (params) => {
  return service({
    url: '/videoTopic/deleteVideoTopicByIds',
    method: 'delete',
    params
  })
}

// @Tags VideoTopic
// @Summary 更新videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoTopic true "更新videoTopic表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoTopic/updateVideoTopic [put]
export const updateVideoTopic = (data) => {
  return service({
    url: '/videoTopic/updateVideoTopic',
    method: 'put',
    data
  })
}

// @Tags VideoTopic
// @Summary 用id查询videoTopic表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.VideoTopic true "用id查询videoTopic表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoTopic/findVideoTopic [get]
export const findVideoTopic = (params) => {
  return service({
    url: '/videoTopic/findVideoTopic',
    method: 'get',
    params
  })
}

// @Tags VideoTopic
// @Summary 分页获取videoTopic表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取videoTopic表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoTopic/getVideoTopicList [get]
export const getVideoTopicList = (params) => {
  return service({
    url: '/videoTopic/getVideoTopicList',
    method: 'get',
    params
  })
}

// @Tags VideoTopic
// @Summary 不需要鉴权的videoTopic表接口
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoTopicSearch true "分页获取videoTopic表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoTopic/getVideoTopicPublic [get]
export const getVideoTopicPublic = () => {
  return service({
    url: '/videoTopic/getVideoTopicPublic',
    method: 'get',
  })
}
