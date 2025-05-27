import service from '@/utils/request'
// @Tags VideoUser
// @Summary 创建用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoUser true "创建用户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /videoUser/createVideoUser [post]
export const createVideoUser = (data) => {
  return service({
    url: '/videoUser/createVideoUser',
    method: 'post',
    data
  })
}

// @Tags VideoUser
// @Summary 删除用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoUser true "删除用户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoUser/deleteVideoUser [delete]
export const deleteVideoUser = (params) => {
  return service({
    url: '/videoUser/deleteVideoUser',
    method: 'delete',
    params
  })
}

// @Tags VideoUser
// @Summary 批量删除用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoUser/deleteVideoUser [delete]
export const deleteVideoUserByIds = (params) => {
  return service({
    url: '/videoUser/deleteVideoUserByIds',
    method: 'delete',
    params
  })
}

// @Tags VideoUser
// @Summary 更新用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.VideoUser true "更新用户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoUser/updateVideoUser [put]
export const updateVideoUser = (data) => {
  return service({
    url: '/videoUser/updateVideoUser',
    method: 'put',
    data
  })
}

// @Tags VideoUser
// @Summary 用id查询用户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.VideoUser true "用id查询用户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoUser/findVideoUser [get]
export const findVideoUser = (params) => {
  return service({
    url: '/videoUser/findVideoUser',
    method: 'get',
    params
  })
}

// @Tags VideoUser
// @Summary 分页获取用户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoUser/getVideoUserList [get]
export const getVideoUserList = (params) => {
  return service({
    url: '/videoUser/getVideoUserList',
    method: 'get',
    params
  })
}

// @Tags VideoUser
// @Summary 不需要鉴权的用户管理接口
// @Accept application/json
// @Produce application/json
// @Param data query videoReq.VideoUserSearch true "分页获取用户管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /videoUser/getVideoUserPublic [get]
export const getVideoUserPublic = () => {
  return service({
    url: '/videoUser/getVideoUserPublic',
    method: 'get',
  })
}
