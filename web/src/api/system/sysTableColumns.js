import service from '@/utils/request'
// @Tags SysTableColumns
// @Summary 创建表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysTableColumns true "创建表字段配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysTableColumns/createSysTableColumns [post]
export const createSysTableColumns = (data) => {
  return service({
    url: '/sysTableColumns/createSysTableColumns',
    method: 'post',
    data
  })
}

// @Tags SysTableColumns
// @Summary 删除表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysTableColumns true "删除表字段配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTableColumns/deleteSysTableColumns [delete]
export const deleteSysTableColumns = (params) => {
  return service({
    url: '/sysTableColumns/deleteSysTableColumns',
    method: 'delete',
    params
  })
}

// @Tags SysTableColumns
// @Summary 批量删除表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除表字段配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysTableColumns/deleteSysTableColumns [delete]
export const deleteSysTableColumnsByIds = (params) => {
  return service({
    url: '/sysTableColumns/deleteSysTableColumnsByIds',
    method: 'delete',
    params
  })
}

// @Tags SysTableColumns
// @Summary 更新表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysTableColumns true "更新表字段配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysTableColumns/updateSysTableColumns [put]
export const updateSysTableColumns = (data) => {
  return service({
    url: '/sysTableColumns/updateSysTableColumns',
    method: 'put',
    data
  })
}

export const updateSysTableColumnsInfo = (data) => {
  return service({
    url: '/sysTableColumns/updateSysTableColumnsInfo',
    method: 'put',
    data
  })
}

// @Tags SysTableColumns
// @Summary 用id查询表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysTableColumns true "用id查询表字段配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysTableColumns/findSysTableColumns [get]
export const findSysTableColumns = (params) => {
  return service({
    url: '/sysTableColumns/findSysTableColumns',
    method: 'get',
    params
  })
}

// @Tags SysTableColumns
// @Summary 分页获取表字段配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取表字段配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysTableColumns/getSysTableColumnsList [get]
export const getSysTableColumnsList = (params) => {
  return service({
    url: '/sysTableColumns/getSysTableColumnsList',
    method: 'get',
    params
  })
}
export const getStructNameColumns = (params) => {
  return service({
    url: '/sysTableColumns/getStructNameColumns',
    method: 'get',
    params
  })
}

// @Tags SysTableColumns
// @Summary 不需要鉴权的表字段配置接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysTableColumnsSearch true "分页获取表字段配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysTableColumns/getSysTableColumnsPublic [get]
export const getSysTableColumnsPublic = () => {
  return service({
    url: '/sysTableColumns/getSysTableColumnsPublic',
    method: 'get',
  })
}

export const syncSysTableColumnsInfo = (params) => {
  return service({
    url: '/sysTableColumns/syncSysTableColumnsInfo',
    method: 'get',
    params
  })
}
