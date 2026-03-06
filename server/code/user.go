package code

// 认证相关错误码 - 按设计规范
var (
	// 基础用户错误 (1001-1099)
	ErrUserNotFound  = NewError(1001, "用户不存在")
	ErrUserLocked    = NewError(1003, "用户已被锁定")
	ErrPasswordError = NewError(1002, "密码错误")

	// TOTP相关错误 (1101-1199)
	ErrTotpNotBound  = NewError(1101, "TOTP未绑定")
	ErrTotpCodeError = NewError(1102, "TOTP验证码错误或过期")

	// Passkey相关错误 (1201-1299)
	ErrPasskeyNotBound = NewError(1201, "通行密钥未绑定")
	ErrPasskeyInvalid  = NewError(1202, "通行密钥验证失败")

	// 绑定相关错误 (1301-1399)
	ErrBindTooFrequent = NewError(1301, "绑定操作过于频繁")

	// 恢复码错误 (1401-1499)
	ErrRecoveryCodeInvalid = NewError(1401, "恢复码无效")

	// 安全相关错误 (1501-1599)
	ErrCSRFFailed = NewError(1501, "CSRF验证失败")

	// 频率限制错误 (1601-1699)
	ErrRateLimited = NewError(1601, "操作过于频繁，请稍后再试")
)

// 用户管理错误 - 保持原有
var (
	ErrUserExists            = NewError(CodeUserError, "用户已存在")
	ErrDeleteSelfNotAllowed  = NewError(CodeUserError, "不能删除自己")
	ErrOldPasswordError      = NewError(CodeUserError, "旧密码错误")
	ErrSetInfoFailed         = NewError(CodeUserError, "设置用户信息失败")
	ErrGetInfoFailed         = NewError(CodeUserError, "获取用户信息失败")
	ErrBindGoogleAuthFailed  = NewError(CodeUserError, "绑定谷歌验证器失败")
	ErrResetGoogleAuthFailed = NewError(CodeUserError, "重置谷歌验证器失败")
	ErrResetPasswordFailed   = NewError(CodeUserError, "重置密码失败")
)
