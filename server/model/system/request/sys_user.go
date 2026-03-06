package request

import (
	"hab/enum"
	common "hab/model/common/request"
	"hab/model/system"
)

// Register User register structure
type Register struct {
	Username     string           `json:"userName" example:"用户名"`
	Password     string           `json:"passWord" example:"密码"`
	NickName     string           `json:"nickName" example:"昵称"`
	HeaderImg    string           `json:"headerImg" example:"头像链接"`
	AuthorityId  uint             `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int              `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint           `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string           `json:"phone" example:"电话号码"`
	Email        string           `json:"email" example:"电子邮箱"`
	Type         enum.SysUserType `json:"type" example:"用户类型"`
	Parameter    string           `json:"parameter" example:"用户参数"`
}

// Login User login structure
type Login struct {
	LoginType      string `json:"loginType"`      // 登录类型
	Username       string `json:"username"`       // 用户名
	Password       string `json:"password"`       // 密码
	Captcha        string `json:"captcha"`        // 验证码
	CaptchaId      string `json:"captchaId"`      // 验证码ID
	GoogleAuthCode string `json:"googleAuthCode"` // 谷歌验证码
	Passkey        string `json:"passkey"`        // Passkey
	Domain         string `json:"domain"`         // 域名
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth Modify user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// SetUserAuthorities Modify user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/HAB_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
	Language     string                `json:"language" gorm:"default:zh-CN;comment:用户语言"` // 用户语言
	Type         enum.SysUserType      `json:"type" gorm:"default:1;comment:用户类型"`         // 用户类型
	Parameter    string                `json:"parameter" gorm:"default:null;comment:用户参数"` // 用户参数
}

type GetUserList struct {
	common.PageInfo
	Username string `json:"username" form:"username"`
	NickName string `json:"nickName" form:"nickName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
}

// BindGoogleAuth 绑定谷歌验证器请求
type BindGoogleAuth struct {
	Code      string `json:"code"`      // 验证码
	SecretKey string `json:"secretKey"` // 密钥
}

// BindPasskey 绑定Passkey请求
type BindPasskey struct {
	PasskeyData string `json:"passkeyData"` // Passkey数据
}

// BindGoogleAuthByLogin 登录前绑定谷歌验证器请求
type BindGoogleAuthByLogin struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Code      string `json:"code"`      // 验证码
	SecretKey string `json:"secretKey"` // 密钥
}

// BindPasskeyByLogin 登录前绑定Passkey请求
type BindPasskeyByLogin struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	PasskeyData string `json:"passkeyData"` // Passkey数据
}

// 新的登录API请求结构体 - 按设计规范

// SecurityStateReq 查询用户安全状态请求
type SecurityStateReq struct {
	Username string `json:"username"` // 用户名
}

// PasswordVerifyReq 密码验证请求
type PasswordVerifyReq struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// TotpLoginReq TOTP登录请求
type TotpLoginReq struct {
	Username string `json:"username"` // 用户名
	Code     string `json:"code"`     // TOTP验证码
}

// PasskeyAssertionOptionsReq Passkey登录选项请求
type PasskeyAssertionOptionsReq struct {
	Username string `json:"username"` // 用户名（可选）
}

// PasskeyAssertionVerifyReq Passkey登录验证请求
type PasskeyAssertionVerifyReq struct {
	Id       string                                `json:"id"`
	RawId    string                                `json:"rawId"`
	Type     string                                `json:"type"`
	Response PasskeyAuthenticatorAssertionResponse `json:"response"`
}

type PasskeyAuthenticatorAssertionResponse struct {
	ClientDataJSON    string `json:"clientDataJSON"`
	AuthenticatorData string `json:"authenticatorData"`
	Signature         string `json:"signature"`
	UserHandle        string `json:"userHandle"`
}

// TotpBindInitReq TOTP绑定初始化请求（暂时为空）
type TotpBindInitReq struct{}

// TotpBindVerifyReq TOTP绑定验证请求
type TotpBindVerifyReq struct {
	Code        string `json:"code"`         // TOTP验证码
	Secret      string `json:"secret"`       // 临时传递的secret，实际应该从session获取
	BindSession string `json:"bind_session"` // 绑定会话token
}

// PasskeyAttestationOptionsReq Passkey绑定选项请求
type PasskeyAttestationOptionsReq struct {
	DisplayName string `json:"displayName"` // 设备显示名称
}

// PasskeyAttestationVerifyReq Passkey绑定验证请求
type PasskeyAttestationVerifyReq struct {
	Id          string                                  `json:"id"`
	RawId       string                                  `json:"rawId"`
	Type        string                                  `json:"type"`
	Response    PasskeyAuthenticatorAttestationResponse `json:"response"`
	BindSession string                                  `json:"bind_session"` // 绑定会话token
}

type PasskeyAuthenticatorAttestationResponse struct {
	ClientDataJSON    string `json:"clientDataJSON"`
	AttestationObject string `json:"attestationObject"`
}
