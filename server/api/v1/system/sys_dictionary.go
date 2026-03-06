package system

import (
	"hab/global"
	"hab/model/common/response"
	"hab/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// CreateSysDictionary
// @Tags      SysDictionary
// @Summary   创建SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysDictionary"
// @Router    /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		global.HAB_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Success", c)
}

// DeleteSysDictionary
// @Tags      SysDictionary
// @Summary   删除SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionary"
// @Router    /sysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.DeleteSysDictionary(dictionary)
	if err != nil {
		global.HAB_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Success", c)
}

// UpdateSysDictionary
// @Tags      SysDictionary
// @Summary   更新SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysDictionary"
// @Router    /sysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.UpdateSysDictionary(&dictionary)
	if err != nil {
		global.HAB_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Success", c)
}

// FindSysDictionary
// @Tags      SysDictionary
// @Summary   用id查询SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysDictionary                                       true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysDictionary"
// @Router    /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		global.HAB_LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("Dictionary not created or not opened", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "ok", c)
}

// GetSysDictionaryList
// @Tags      SysDictionary
// @Summary   分页获取SysDictionary列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	list, err := dictionaryService.GetSysDictionaryInfoList()
	if err != nil {
		global.HAB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(list, "Success", c)
}
