package {{.Package}}

import (
    "forge-basic/code"
    "forge-basic/enum"
	"forge-basic/model/common/request"
	{{if not .OnlyTemplate}}
	"{{.Module}}/global"
    "{{.Module}}/model/common/response"
    "{{.Module}}/model/{{.Package}}"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    {{- if .AutoCreateResource}}
    "{{.Module}}/utils"
    {{- end }}
    {{- else}}
    "{{.Module}}/model/common/response"
    "github.com/gin-gonic/gin"
    {{- end}}
)

type {{.StructName}}Api struct {}

{{if not .OnlyTemplate}}

// Create{{.StructName}} 创建{{.Description}}
// @Tags {{.StructName}}
// @Summary 创建{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "创建{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    {{.Abbreviation}}.CreatedBy = utils.GetUserID(c)
	{{- end }}
	err = {{.Abbreviation}}Service.Create{{.StructName}}(&{{.Abbreviation}}, c)
	if err != nil {
        global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Create, err.Error(), c)
		return
	}
    response.OkWithMessage("Create successful", c)
}

// Import{{.StructName}} 导入{{.Description}}
// @Tags {{.StructName}}
// @Summary 导入{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body []{{.Package}}.{{.StructName}} true "导入{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /{{.Abbreviation}}/import{{.StructName}} [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Import{{.StructName}}(c *gin.Context) {
    var data struct {
    		Type enum.ImportType       `json:"type" form:"type"`
    		List []{{.Package}}.{{.StructName}} `json:"list" form:"list"`
    	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
	userID := utils.GetUserID(c)
	for i := range data.List {
		data.List[i].CreatedBy = userID
	}
	{{- end }}
	err = {{.Abbreviation}}Service.Import{{.StructName}}(data.Type, data.List, c)
	if err != nil {
		global.GVA_LOG.Error("Import failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Import, err.Error(), c)
		return
	}
	response.OkWithMessage("Import successful", c)
}

// Delete{{.StructName}} 删除{{.Description}}
// @Tags {{.StructName}}
// @Summary 删除{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "删除{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
		{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
        {{- end }}
	err := {{.Abbreviation}}Service.Delete{{.StructName}}({{.PrimaryField.FieldJson}} {{- if .AutoCreateResource -}},userID{{- end -}}, c)
	if err != nil {
        global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Delete, err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}
// @Tags {{.StructName}}
// @Summary 批量删除{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
	{{.PrimaryField.FieldJson}}s := c.QueryArray("{{.PrimaryField.FieldJson}}s[]")
    	{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
        {{- end }}
	err := {{.Abbreviation}}Service.Delete{{.StructName}}ByIds({{.PrimaryField.FieldJson}}s{{- if .AutoCreateResource }},userID{{- end }}, c)
	if err != nil {
        global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Delete, err.Error(), c)
		return
	}
	response.OkWithMessage("Batch delete successful", c)
}

// Update{{.StructName}} 更新{{.Description}}
// @Tags {{.StructName}}
// @Summary 更新{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "更新{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	    {{- if .AutoCreateResource }}
    {{.Abbreviation}}.UpdatedBy = utils.GetUserID(c)
        {{- end }}
	err = {{.Abbreviation}}Service.Update{{.StructName}}({{.Abbreviation}}, c)
	if err != nil {
        global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Update, err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// Find{{.StructName}} 用id查询{{.Description}}
// @Tags {{.StructName}}
// @Summary 用id查询{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param {{.PrimaryField.FieldJson}} query {{.PrimaryField.FieldType}} true "用id查询{{.Description}}"
// @Success 200 {object} response.Response{data={{.Package}}.{{.StructName}},msg=string} "ok"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
	re{{.Abbreviation}}, err := {{.Abbreviation}}Service.Get{{.StructName}}({{.PrimaryField.FieldJson}}, c)
	if err != nil {
        global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Get, err.Error(), c)
		return
	}
	response.OkWithData(re{{.Abbreviation}}, c)
}

{{- if .IsTree }}
// Get{{.StructName}}List 分页获取{{.Description}}列表,Tree模式下不接受参数
// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Success"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	list, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(, c)
	if err != nil {
	    global.GVA_LOG.Error("Get list failed!", zap.Error(err))
        response.FailWithErrData(code.Err_Business_Get, err.Error(), c)
        return
    }
    response.OkWithDetailed(list, "Get list successful", c)
}
{{- else }}
// Get{{.StructName}}List 分页获取{{.Description}}列表
// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Success"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	var pageInfo request.QueryInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	if pageInfo.PageInfo == nil || pageInfo.PageInfo.Page <= 0 || pageInfo.PageInfo.PageSize <= 0 {
		pageInfo.PageInfo = &request.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}

	if pageInfo.SortInfo == nil || len(pageInfo.SortInfo) == 0 {
		pageInfo.SortInfo = map[string]*request.SortInfo{
			"id": &request.SortInfo{
				Index: 0,
				Type:  "desc",
			},
		}
	}

	list, total, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(pageInfo, c)
	if err != nil {
	    global.GVA_LOG.Error("Get list failed!", zap.Error(err))
        response.FailWithErrData(code.Err_Business_List,err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.PageInfo.Page,
        PageSize: pageInfo.PageInfo.PageSize,
    }, "Get list successful", c)
}
{{- end }}

{{- if .HasDataSource }}
// Get{{.StructName}}DataSource 获取{{.StructName}}的数据源
// @Tags {{.StructName}}
// @Summary 获取{{.StructName}}的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "ok"
// @Router /{{.Abbreviation}}/get{{.StructName}}DataSource [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}DataSource(c *gin.Context) {
    // 此接口为获取数据源定义的数据
    dataSource, err := {{.Abbreviation}}Service.Get{{.StructName}}DataSource()
    if err != nil {
        global.GVA_LOG.Error("Query failed!", zap.Error(err))
   		response.FailWithErrData(code.Err_Business_List,err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}
{{- end }}

{{- end }}

// Get{{.StructName}}Public 不需要鉴权的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 不需要鉴权的{{.Description}}接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /{{.Abbreviation}}/get{{.StructName}}Public [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}Public(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    {{.Abbreviation}}Service.Get{{.StructName}}Public()
    response.OkWithDetailed(gin.H{
       "info": "Public {{.Description}} interface information",
    }, "Get successful", c)
}
