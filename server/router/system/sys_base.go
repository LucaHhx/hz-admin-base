package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		// 保留兼容接口
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}

	// 新的认证路由组 - 按设计规范
	authRouter := Router.Group("auth")
	{
		// 登录模式查询（公开接口）
		authRouter.GET("login-mode", baseApi.GetLoginMode)

		// 用户安全状态查询
		authRouter.POST("security-state", baseApi.SecurityState)

		// 密码验证
		authRouter.POST("password/verify", baseApi.PasswordVerify)
		authRouter.POST("password/login", baseApi.PasswordLogin)

		// TOTP登录和绑定
		authRouter.POST("totp/login", baseApi.TotpLogin)
		authRouter.POST("totp/bind/init", baseApi.TotpBindInit)
		authRouter.POST("totp/bind/verify", baseApi.TotpBindVerify)

		// Passkey登录和绑定
		authRouter.POST("passkey/assertion/options", baseApi.PasskeyAssertionOptions)
		authRouter.POST("passkey/assertion/verify", baseApi.PasskeyAssertionVerify)
		authRouter.POST("passkey/attestation/options", baseApi.PasskeyAttestationOptions)
		authRouter.POST("passkey/attestation/verify", baseApi.PasskeyAttestationVerify)
	}

	return baseRouter
}
