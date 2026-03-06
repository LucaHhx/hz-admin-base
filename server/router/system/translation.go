package system

import (
	"hab/api/v1/system"

	"github.com/gin-gonic/gin"
)

type TranslationRouter struct{}

func (s *TranslationRouter) InitTranslationRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	translationRouter := Router.Group("translation")
	translationRouterPublic := RouterPublic.Group("translation")
	translationApi := system.TranslationApi{}
	{
		translationRouter.GET("tree", translationApi.GetFileTree)
		translationRouter.GET("file", translationApi.GetFileContent)
		translationRouter.POST("update", translationApi.UpdateTranslation)
		translationRouter.POST("copy", translationApi.CopyLanguage)
		translationRouter.POST("batch", translationApi.BatchUpdate)
		translationRouter.POST("export-language", translationApi.ExportLanguage)
		translationRouter.POST("delete-language", translationApi.DeleteLanguage)
		translationRouter.GET("download-language", translationApi.DownloadLanguage)
		translationRouter.POST("upload-language", translationApi.UploadLanguage)
	}
	{
		translationRouterPublic.GET("messages", translationApi.GetMessages)
	}
}
