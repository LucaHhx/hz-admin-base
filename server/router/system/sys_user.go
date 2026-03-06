package system

import (
	"hz-admin-base/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup, pubRouter *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := pubRouter.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
		userRouter.PUT("setSelfSetting", baseApi.SetSelfSetting)          // 用户界面配置
		userRouter.GET("getGoogleAuthQR", baseApi.GetGoogleAuthQR)        // 获取谷歌验证器二维码
		//userRouter.POST("bindGoogleAuth", baseApi.BindGoogleAuth)         // 绑定谷歌验证器
		userRouter.POST("resetGoogleAuth", baseApi.ResetGoogleAuth) // 重置谷歌验证器
		//userRouter.POST("bindPasskey", baseApi.BindPasskey)         // 绑定Passkey
		userRouter.POST("unbindSecurity", baseApi.UnbindSecurity) // 解绑安全验证
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)                     // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)                      // 获取自身信息
		userRouterWithoutRecord.POST("bindGoogleAuthByLogin", baseApi.BindGoogleAuthByLogin) // 登录前绑定谷歌验证器
		userRouterWithoutRecord.POST("bindPasskeyByLogin", baseApi.BindPasskeyByLogin)       // 登录前绑定Passkey
	}
}
