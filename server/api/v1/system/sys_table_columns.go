package system

import (
	"hz-admin-base/global"
	"hz-admin-base/model/common/response"
	"hz-admin-base/model/system"
	systemReq "hz-admin-base/model/system/request"
	"hz-admin-base/utils"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysTableColumnsApi struct{}

// CreateSysTableColumns 创建表字段配置
// @Tags SysTableColumns
// @Summary 创建表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysTableColumns true "创建表字段配置"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysTableColumns/createSysTableColumns [post]
func (sysTableColumnsApi *SysTableColumnsApi) CreateSysTableColumns(c *gin.Context) {
	var sysTableColumns system.SysTableColumns
	err := c.ShouldBindJSON(&sysTableColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysTableColumns.CreatedBy = utils.GetUserID(c)
	err = sysTableColumnsService.CreateSysTableColumns(&sysTableColumns)
	if err != nil {
		global.GVA_LOG.Error("Create failed!", zap.Error(err))
		response.FailWithMessage("Create failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Clear()
	response.OkWithMessage("Create successful", c)
}

// ImportSysTableColumns 导入表字段配置
// @Tags SysTableColumns
// @Summary 导入表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body []system.SysTableColumns true "导入表字段配置"
// @Success 200 {object} response.Response{msg=string} "导入成功"
// @Router /sysTableColumns/importSysTableColumns [post]
func (sysTableColumnsApi *SysTableColumnsApi) ImportSysTableColumns(c *gin.Context) {
	var sysTableColumnsList []system.SysTableColumns
	err := c.ShouldBindJSON(&sysTableColumnsList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	for i := range sysTableColumnsList {
		sysTableColumnsList[i].CreatedBy = userID
	}
	err = sysTableColumnsService.ImportSysTableColumns(sysTableColumnsList)
	if err != nil {
		global.GVA_LOG.Error("Import failed!", zap.Error(err))
		response.FailWithMessage("Import failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Clear()
	response.OkWithMessage("Import successful", c)
}

// DeleteSysTableColumns 删除表字段配置
// @Tags SysTableColumns
// @Summary 删除表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysTableColumns true "删除表字段配置"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysTableColumns/deleteSysTableColumns [delete]
func (sysTableColumnsApi *SysTableColumnsApi) DeleteSysTableColumns(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := sysTableColumnsService.DeleteSysTableColumns(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("Delete failed!", zap.Error(err))
		response.FailWithMessage("Delete failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Clear()
	response.OkWithMessage("Delete successful", c)
}

// DeleteSysTableColumnsByIds 批量删除表字段配置
// @Tags SysTableColumns
// @Summary 批量删除表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysTableColumns/deleteSysTableColumnsByIds [delete]
func (sysTableColumnsApi *SysTableColumnsApi) DeleteSysTableColumnsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := sysTableColumnsService.DeleteSysTableColumnsByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("Batch delete failed!", zap.Error(err))
		response.FailWithMessage("Batch delete failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Clear()
	response.OkWithMessage("Batch delete successful", c)
}

// UpdateSysTableColumns 更新表字段配置
// @Tags SysTableColumns
// @Summary 更新表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysTableColumns true "更新表字段配置"
// @Success 200 {object} response.Response{msg=string} "Success"
// @Router /sysTableColumns/updateSysTableColumns [put]
func (sysTableColumnsApi *SysTableColumnsApi) UpdateSysTableColumns(c *gin.Context) {
	var sysTableColumns system.SysTableColumns
	err := c.ShouldBindJSON(&sysTableColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysTableColumns.UpdatedBy = utils.GetUserID(c)
	err = sysTableColumnsService.UpdateSysTableColumns(sysTableColumns)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Clear()
	response.OkWithMessage("Update successful", c)
}

// FindSysTableColumns 用id查询表字段配置
// @Tags SysTableColumns
// @Summary 用id查询表字段配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询表字段配置"
// @Success 200 {object} response.Response{data=system.SysTableColumns,msg=string} "ok"
// @Router /sysTableColumns/findSysTableColumns [get]
func (sysTableColumnsApi *SysTableColumnsApi) FindSysTableColumns(c *gin.Context) {
	ID := c.Query("ID")
	resysTableColumns, err := sysTableColumnsService.GetSysTableColumns(ID)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed: "+err.Error(), c)
		return
	}
	response.OkWithData(resysTableColumns, c)
}

// GetSysTableColumnsList 分页获取表字段配置列表
// @Tags SysTableColumns
// @Summary 分页获取表字段配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysTableColumnsSearch true "分页获取表字段配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Success"
// @Router /sysTableColumns/getSysTableColumnsList [get]
func (sysTableColumnsApi *SysTableColumnsApi) GetSysTableColumnsList(c *gin.Context) {
	var pageInfo systemReq.SysTableColumnsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysTableColumnsService.GetSysTableColumnsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Get list failed!", zap.Error(err))
		response.FailWithMessage("Get list failed: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get list successful", c)
}

func (sysTableColumnsApi *SysTableColumnsApi) GetStructNameColumns(c *gin.Context) {
	structName := c.Query("structName")
	columns, err := sysTableColumnsService.GetStructNameColumns(structName)
	if err != nil {
		global.GVA_LOG.Error("Get structName columns failed!", zap.Error(err))
		response.FailWithMessage("Get structName columns failed: "+err.Error(), c)
		return
	}
	response.OkWithData(columns, c)

}

// GetSysTableColumnsPublic 不需要鉴权的表字段配置接口
// @Tags SysTableColumns
// @Summary 不需要鉴权的表字段配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /sysTableColumns/getSysTableColumnsPublic [get]
func (sysTableColumnsApi *SysTableColumnsApi) GetSysTableColumnsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysTableColumnsService.GetSysTableColumnsPublic()
	response.OkWithDetailed(gin.H{
		"info": "Public 表字段配置 interface information",
	}, "Get successful", c)
}

func (sysTableColumnsApi *SysTableColumnsApi) UpdateSysTableColumnsInfo(c *gin.Context) {
	var sysTableColumnsList []system.SysTableColumns
	err := c.ShouldBindJSON(&sysTableColumnsList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	updatedBy := utils.GetUserID(c)
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range sysTableColumnsList {
			err = tx.Model(&system.SysTableColumns{}).
				Where("id = ?", item.ID).
				Updates(map[string]interface{}{
					"with":             item.With,
					"fixed":            item.Fixed,
					"updated_by":       updatedBy,
					"sort":             item.Sort,
					"format":           item.Format,
					"type":             item.Type,
					"enum":             item.Enum,
					"note":             item.Note,
					"default_filter":   item.DefaultFilter,
					"filter_list":      item.FilterList,
					"filter":           item.Filter,
					"sortable":         item.Sortable,
					"is_add_search":    item.IsAddSearch,
					"search_width":     item.SearchWidth,
					"is_additional":    item.IsAdditional,
					"is_show":          item.IsShow,
					"filter_id":        item.FilterId,
					"form_with":        item.FormWith,
					"form_disabled":    item.FormDisabled,
					"form_hidden":      item.FormHidden,
					"form_order":       item.FormOrder,
					"form_must":        item.FormMust,
					"editor_filter_id": item.EditorFilterId,
					"is_sum":           item.IsSum,
				}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed: "+err.Error(), c)
		return
	}
	utils.ColumnsCache.Remove(utils.GetUserAuthorityId(c))
	response.OkWithMessage("Update successful", c)
}

func (sysTableColumnsApi *SysTableColumnsApi) SyncSysTableColumnsInfo(c *gin.Context) {
	name := c.Query("name")
	tbaleInfoList := utils.GetTableInfo(name).Fields
	var sysTableColumnsList []system.SysTableColumns
	err := global.GVA_DB.Find(&sysTableColumnsList, "struct_name=?", name).Error
	if err != nil {
		global.GVA_LOG.Error("Get table columns failed!", zap.Error(err))
		response.FailWithMessage("Get table columns failed: "+err.Error(), c)
		return
	}
	if len(sysTableColumnsList) == 0 {
		response.FailWithMessage("No table columns found for struct: "+name, c)
		return
	}
	infoTem := sysTableColumnsList[0]
	addList := make([]system.SysTableColumns, 0)
	for _, info := range tbaleInfoList {
		if _, ok := lo.Find(sysTableColumnsList, func(item system.SysTableColumns) bool {
			return item.JsonName == info.JSONTag
		}); ok {
			continue
		}
		jsonName := info.JSONTag
		if jsonName == "" {
			jsonName = utils.LowerFirst(info.Name)
		}
		addList = append(addList, system.SysTableColumns{
			TbName:        infoTem.TbName,
			MenuId:        infoTem.MenuId,
			StructName:    infoTem.StructName,
			JsonName:      info.JSONTag,
			ColumnName:    info.ColumnTag,
			With:          100,
			Type:          info.Type,
			DefaultFilter: "=",
			Sort:          99,
			Note:          info.Comment,
			IsShow:        true,
			Format:        "",
			IsAddSearch:   false,
			SearchWidth:   200,
			FormInfo: system.FormInfo{
				FormWith:  50,
				FormOrder: 99,
			},
		})
	}

	delList := make([]system.SysTableColumns, 0)
	for _, info := range sysTableColumnsList {
		if _, ok := lo.Find(tbaleInfoList, func(item utils.FieldInfo) bool {
			return item.JSONTag == info.JsonName
		}); !ok {
			delList = append(delList, info)
		}
	}
	response.OkWithData(gin.H{
		"addList": addList,
		"delList": delList,
	}, c)
}
