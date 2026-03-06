package response

import (
	"hab/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

// GoogleAuthResponse 谷歌验证器响应
type GoogleAuthResponse struct {
	QRCode    string `json:"qrCode"`    // 二维码内容
	SecretKey string `json:"secretKey"` // 密钥
}

// 新的登录API响应结构体 - 按设计规范

// SecurityStateRes 用户安全状态响应
type SecurityStateRes struct {
	Exists      bool   `json:"exists"`       // 用户是否存在
	HasTotp     bool   `json:"has_totp"`     // 是否已绑定TOTP
	HasPasskey  bool   `json:"has_passkey"`  // 是否已绑定Passkey
	IsLocked    bool   `json:"is_locked"`    // 是否被锁定
	DisplayName string `json:"display_name"` // 显示名称
}

// PasswordVerifyRes 密码验证响应
type PasswordVerifyRes struct {
	Ok          bool   `json:"ok"`           // 验证是否成功
	Next        string `json:"next"`         // 下一步操作
	BindSession string `json:"bind_session"` // 绑定会话token
}

// TotpBindInitRes TOTP绑定初始化响应
type TotpBindInitRes struct {
	QrcodeUrl    string `json:"qrcode_url"`    // 二维码URL（前端用此生成二维码）
	OtpauthUrl   string `json:"otpauth_url"`   // otpauth URL
	SecretMasked string `json:"secret_masked"` // 掩码后的密钥
	Secret       string `json:"secret"`        // 完整密钥（临时返回，实际应该存在session中）
}

// TotpBindVerifyRes TOTP绑定验证响应
type TotpBindVerifyRes struct {
	Bound bool `json:"bound"` // 是否绑定成功
}

// PasskeyAssertionOptionsRes Passkey登录选项响应
type PasskeyAssertionOptionsRes struct {
	PublicKey PasskeyPublicKeyCredentialRequestOptions `json:"publicKey"`
}

// PasskeyAttestationOptionsRes Passkey绑定选项响应
type PasskeyAttestationOptionsRes struct {
	PublicKey PasskeyPublicKeyCredentialCreationOptions `json:"publicKey"`
}

// PasskeyAttestationVerifyRes Passkey绑定验证响应
type PasskeyAttestationVerifyRes struct {
	Bound bool `json:"bound"` // 是否绑定成功
}

// Passkey相关结构体
type PasskeyPublicKeyCredentialRequestOptions struct {
	Challenge        string                                 `json:"challenge"`
	RpId             string                                 `json:"rpId"`
	AllowCredentials []PasskeyPublicKeyCredentialDescriptor `json:"allowCredentials"`
	UserVerification string                                 `json:"userVerification"`
	Timeout          int                                    `json:"timeout"`
}

type PasskeyPublicKeyCredentialDescriptor struct {
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	Transports []string `json:"transports"`
}

type PasskeyPublicKeyCredentialCreationOptions struct {
	Challenge              string                        `json:"challenge"`
	Rp                     PasskeyRelyingParty           `json:"rp"`
	User                   PasskeyUser                   `json:"user"`
	PubKeyCredParams       []PasskeyPubKeyCredParam      `json:"pubKeyCredParams"`
	Timeout                int                           `json:"timeout"`
	Attestation            string                        `json:"attestation"`
	AuthenticatorSelection PasskeyAuthenticatorSelection `json:"authenticatorSelection"`
}

type PasskeyRelyingParty struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type PasskeyUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type PasskeyPubKeyCredParam struct {
	Type string `json:"type"`
	Alg  int    `json:"alg"`
}

type PasskeyAuthenticatorSelection struct {
	ResidentKey      string `json:"residentKey"`
	UserVerification string `json:"userVerification"`
}
