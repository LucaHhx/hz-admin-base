import service from '@/utils/request'

{{- if not .OnlyTemplate}}
// @Tags {{.StructName}}
// @Summary 创建{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "创建{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
export const create{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/create{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 删除{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "删除{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}',
    method: 'delete',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 批量删除{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
export const delete{{.StructName}}ByIds = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}ByIds',
    method: 'delete',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 更新{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "更新{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
export const update{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/update{{.StructName}}',
    method: 'put',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 用id查询{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.{{.StructName}} true "用id查询{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
export const find{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/find{{.StructName}}',
    method: 'get',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取{{.Description}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
export const get{{.StructName}}List = (data) => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}List',
    method: 'post',
    data
  })
}

{{- if .HasDataSource}}
// @Tags {{.StructName}}
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /{{.Abbreviation}}/find{{.StructName}}DataSource [get]
export const get{{.StructName}}DataSource = () => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}DataSource',
    method: 'get',
  })
}
{{- end}}

{{- end}}

// @Tags {{.StructName}}
// @Summary 不需要鉴权的{{.Description}}接口
// @Accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /{{.Abbreviation}}/get{{.StructName}}Public [get]
export const get{{.StructName}}Public = () => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}Public',
    method: 'get',
  })
}

// @Tags {{.StructName}}
// @Summary 导入{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.Import{{.StructName}}Req true "导入{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
export const import{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/import{{.StructName}}',
    method: 'post',
    data
  })
}