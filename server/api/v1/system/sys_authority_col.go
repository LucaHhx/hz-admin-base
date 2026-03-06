package system

import (
	"hz-admin-base/global"
	"hz-admin-base/model/common/response"
	"hz-admin-base/model/system/request"
	"hz-admin-base/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityColApi struct{}

// GetAuthorityCol
// @Tags      AuthorityCol
// @Summary   获取角色列表权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body      request.SysAuthorityColReq                                          true  "菜单ID，角色ID"
// @Success   200  {object}  response.Response{data=systemRes.SysAuthorityColRes,msg=string}  "返回列权限列表"
// @Router    /authorityCol/getAuthorityCol [post]
func (a *AuthorityColApi) GetAuthorityCol(c *gin.Context) {
	var req request.SysAuthorityColReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := authorityColService.GetAuthorityCol(req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(res, "Success", c)
}

// SetAuthorityCol
// @Tags      AuthorityCol
// @Summary   设置角色列表权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body      request.SysAuthorityColReq     true  "菜单ID，角色ID，列ID们（批量设置）"
// @Success   200  {object}  response.Response{msg=string}  "设置角色列表权限"
// @Router    /authorityCol/setAuthorityCol [post]
func (a *AuthorityColApi) SetAuthorityCol(c *gin.Context) {
	var req request.SysAuthorityColReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityColService.SetAuthorityCol(req)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	utils.ColumnsCache.Remove(req.AuthorityId)
	response.OkWithMessage("设置成功", c)
}

// CanRemoveAuthorityCol
// @Tags      AuthorityCol
// @Summary   设置列权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "Success"
// @Router    /authorityCol/canRemoveAuthorityCol [post]
func (a *AuthorityColApi) CanRemoveAuthorityCol(c *gin.Context) {
	id := c.Query("id")
	err := authorityColService.CanRemoveAuthorityCol(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Success", c)
}

// GetMenuColumns
// @Tags      AuthorityCol
// @Summary   获取菜单对应的所有列
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     menuName query string true "菜单名称"
// @Success   200  {object}  response.Response{data=[]system.SysTableColumns,msg=string}  "返回菜单的所有列"
// @Router    /authorityCol/getMenuColumns [get]
func (a *AuthorityColApi) GetMenuColumns(c *gin.Context) {
	menuName := c.Query("menuName")
	if menuName == "" {
		response.FailWithMessage("菜单名称不能为空", c)
		return
	}
	columns, err := authorityColService.GetMenuColumns(menuName)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(columns, "Success", c)
}
