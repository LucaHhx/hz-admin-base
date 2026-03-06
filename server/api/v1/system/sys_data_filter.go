package system

import (
	"hz-admin-base/code"
	"hz-admin-base/enum"
	"hz-admin-base/model/common/request"

	"hz-admin-base/global"
	"hz-admin-base/model/common/response"
	"hz-admin-base/model/system"
	"hz-admin-base/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDataFilterApi struct{}

// CreateSysDataFilter 创建sysDataFilter
// @Tags SysDataFilter
// @Summary 创建sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysDataFilter true "创建sysDataFilter"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysDataFilter/createSysDataFilter [post]
func (sysDataFilterApi *SysDataFilterApi) CreateSysDataFilter(c *gin.Context) {
	var sysDataFilter system.SysDataFilter
	err := c.ShouldBindJSON(&sysDataFilter)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	sysDataFilter.CreatedBy = utils.GetUserID(c)
	err = sysDataFilterService.CreateSysDataFilter(&sysDataFilter, c)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Create, err.Error(), c)
		return
	}
	response.OkWithMessage("Create successful", c)
}

// ImportSysDataFilter 导入sysDataFilter
// @Tags SysDataFilter
// @Summary 导入sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body []system.SysDataFilter true "导入sysDataFilter"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /sysDataFilter/importSysDataFilter [post]
func (sysDataFilterApi *SysDataFilterApi) ImportSysDataFilter(c *gin.Context) {
	var data struct {
		Type enum.ImportType        `json:"type" form:"type"`
		List []system.SysDataFilter `json:"list" form:"list"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	for i := range data.List {
		data.List[i].CreatedBy = userID
	}
	err = sysDataFilterService.ImportSysDataFilter(data.Type, data.List, c)
	if err != nil {
		global.GVA_LOG.Error("Import failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Import, err.Error(), c)
		return
	}
	response.OkWithMessage("Import successful", c)
}

// DeleteSysDataFilter 删除sysDataFilter
// @Tags SysDataFilter
// @Summary 删除sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysDataFilter true "删除sysDataFilter"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysDataFilter/deleteSysDataFilter [delete]
func (sysDataFilterApi *SysDataFilterApi) DeleteSysDataFilter(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := sysDataFilterService.DeleteSysDataFilter(ID, userID, c)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Delete, err.Error(), c)
		return
	}
	response.OkWithMessage("Delete successful", c)
}

// DeleteSysDataFilterByIds 批量删除sysDataFilter
// @Tags SysDataFilter
// @Summary 批量删除sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysDataFilter/deleteSysDataFilterByIds [delete]
func (sysDataFilterApi *SysDataFilterApi) DeleteSysDataFilterByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := sysDataFilterService.DeleteSysDataFilterByIds(IDs, userID, c)
	if err != nil {
		global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Delete, err.Error(), c)
		return
	}
	response.OkWithMessage("Batch delete successful", c)
}

// UpdateSysDataFilter 更新sysDataFilter
// @Tags SysDataFilter
// @Summary 更新sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysDataFilter true "更新sysDataFilter"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysDataFilter/updateSysDataFilter [put]
func (sysDataFilterApi *SysDataFilterApi) UpdateSysDataFilter(c *gin.Context) {
	var sysDataFilter system.SysDataFilter
	err := c.ShouldBindJSON(&sysDataFilter)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	sysDataFilter.UpdatedBy = utils.GetUserID(c)
	err = sysDataFilterService.UpdateSysDataFilter(sysDataFilter, c)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Update, err.Error(), c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindSysDataFilter 用id查询sysDataFilter
// @Tags SysDataFilter
// @Summary 用id查询sysDataFilter
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询sysDataFilter"
// @Success 200 {object} response.Response{data=system.SysDataFilter,msg=string} "ok"
// @Router /sysDataFilter/findSysDataFilter [get]
func (sysDataFilterApi *SysDataFilterApi) FindSysDataFilter(c *gin.Context) {
	ID := c.Query("ID")
	resysDataFilter, err := sysDataFilterService.GetSysDataFilter(ID, c)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Get, err.Error(), c)
		return
	}
	response.OkWithData(resysDataFilter, c)
}

// GetSysDataFilterList 分页获取sysDataFilter列表
// @Tags SysDataFilter
// @Summary 分页获取sysDataFilter列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysDataFilterSearch true "分页获取sysDataFilter列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Success"
// @Router /sysDataFilter/getSysDataFilterList [get]
func (sysDataFilterApi *SysDataFilterApi) GetSysDataFilterList(c *gin.Context) {
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

	list, total, err := sysDataFilterService.GetSysDataFilterInfoList(pageInfo, c)
	if err != nil {
		global.GVA_LOG.Error("Get list failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_List, err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.PageInfo.Page,
		PageSize: pageInfo.PageInfo.PageSize,
	}, "Get list successful", c)
}

// GetSysDataFilterPublic 不需要鉴权的sysDataFilter接口
// @Tags SysDataFilter
// @Summary 不需要鉴权的sysDataFilter接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /sysDataFilter/getSysDataFilterPublic [get]
func (sysDataFilterApi *SysDataFilterApi) GetSysDataFilterPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysDataFilterService.GetSysDataFilterPublic()
	response.OkWithDetailed(gin.H{
		"info": "Public sysDataFilter interface information",
	}, "Get successful", c)
}

func (sysDataFilterApi *SysDataFilterApi) ExecuteSql(c *gin.Context) {
	sql := c.Query("sql")
	data, err := sysDataFilterService.ExecuteSql(sql)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Get, err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

func (sysDataFilterApi *SysDataFilterApi) FilterData(c *gin.Context) {
	var data struct {
		Filters []string `json:"filters" form:"filters"`
		Id      uint     `json:"id" form:"id"`
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithErrData(code.Err_Common_InvalidParams, err.Error(), c)
		return
	}
	FilterData, err := sysDataFilterService.FilterData(data.Filters, data.Id)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithErrData(code.Err_Business_Get, err.Error(), c)
		return
	}
	response.OkWithData(FilterData, c)
}
