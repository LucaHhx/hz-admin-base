import service from '@/utils/request'
// @Tags SysDataFilter
// @Summary 创建sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysDataFilter true "创建sysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/createSysDataFilter [post]
export const createSysDataFilter = (data) => {
  return service({
    url: '/sysDataFilter/createSysDataFilter',
    method: 'post',
    data
  })
}

// @Tags SysDataFilter
// @Summary 删除sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysDataFilter true "删除sysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/deleteSysDataFilter [delete]
export const deleteSysDataFilter = (params) => {
  return service({
    url: '/sysDataFilter/deleteSysDataFilter',
    method: 'delete',
    params
  })
}

// @Tags SysDataFilter
// @Summary 批量删除sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/deleteSysDataFilter [delete]
export const deleteSysDataFilterByIds = (params) => {
  return service({
    url: '/sysDataFilter/deleteSysDataFilterByIds',
    method: 'delete',
    params
  })
}

// @Tags SysDataFilter
// @Summary 更新sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysDataFilter true "更新sysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/updateSysDataFilter [put]
export const updateSysDataFilter = (data) => {
  return service({
    url: '/sysDataFilter/updateSysDataFilter',
    method: 'put',
    data
  })
}

// @Tags SysDataFilter
// @Summary 用id查询sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysDataFilter true "用id查询sysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /sysDataFilter/findSysDataFilter [get]
export const findSysDataFilter = (params) => {
  return service({
    url: '/sysDataFilter/findSysDataFilter',
    method: 'get',
    params
  })
}

// @Tags SysDataFilter
// @Summary 生成SQL
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysDataFilter true "生成SQL"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/generateSql [get]
export const executeSql = (params) => {
  return service({
    url: '/sysDataFilter/executeSql',
    method: 'get',
    params
  })
}

// @Tags SysDataFilter
// @Summary 分页获取sysDataFilter列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysDataFilter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /sysDataFilter/getSysDataFilterList [get]
export const getSysDataFilterList = (data) => {
  return service({
    url: '/sysDataFilter/getSysDataFilterList',
    method: 'post',
    data
  })
}

export const filterData = (data) => {
  return service({
    url: '/sysDataFilter/filterData',
    method: 'post',
    data
  })
}

// @Tags SysDataFilter
// @Summary 获取sysDataFilter数据
// @Tags SysDataFilter
// @Summary 不需要鉴权的sysDataFilter接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysDataFilterSearch true "分页获取sysDataFilter列表"
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /sysDataFilter/getSysDataFilterPublic [get]
export const getSysDataFilterPublic = () => {
  return service({
    url: '/sysDataFilter/getSysDataFilterPublic',
    method: 'get',
  })
}

// @Tags SysDataFilter
// @Summary 导入sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.ImportSysDataFilterReq true "导入SysDataFilter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
export const importSysDataFilter = (data) => {
  return service({
    url: '/sysDataFilter/importSysDataFilter',
    method: 'post',
    data
  })
}
