package system

import (
	"errors"
	"hab/enum"
	"hab/utils/pwd"
	"strings"
	"time"

	"hab/code"
	"hab/global"
	"hab/model/common/response"
	"hab/model/system"
	systemReq "hab/model/system/request"
	systemRes "hab/model/system/response"
	"hab/utils"

	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 按照设计规范重写登录系统

// GetLoginMode 获取当前登录模式（公开接口，无需认证）
// @Tags     Auth
// @Summary  获取登录模式
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]string,msg=string}  "登录模式"
// @Router   /auth/login-mode [get]
func (b *BaseApi) GetLoginMode(c *gin.Context) {
	mode := global.HAB_CONFIG.System.LoginMode
	if mode == "" {
		mode = "simple"
	}
	response.OkWithDetailed(gin.H{"login_mode": mode}, "Success", c)
}

// generateChallenge 生成WebAuthn挑战
func generateChallenge() string {
	challenge := make([]byte, 32)
	_, err := rand.Read(challenge)
	if err != nil {
		// 如果随机数生成失败，使用时间戳作为后备
		timeBytes := []byte(fmt.Sprintf("%d", time.Now().UnixNano()))
		if len(timeBytes) < 32 {
			paddedBytes := make([]byte, 32)
			copy(paddedBytes, timeBytes)
			challenge = paddedBytes
		} else {
			challenge = timeBytes[:32]
		}
	}
	return base64.RawURLEncoding.EncodeToString(challenge)
}

// SecurityState 查询用户安全状态（设计规范 5.2.1）
// @Tags     Auth
// @Summary  查询用户安全状态
// @Produce  application/json
// @Param    data  body      systemReq.SecurityStateReq                                    true  "用户名"
// @Success  200   {object}  response.Response{data=systemRes.SecurityStateRes,msg=string}  "用户安全状态"
// @Router   /auth/security-state [post]
func (b *BaseApi) SecurityState(c *gin.Context) {
	var req systemReq.SecurityStateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Username == "" {
		response.FailWithMessage("用户名不能为空", c)
		return
	}

	user, err := userService.GetUserByUsername(req.Username)
	if err != nil {
		response.FailWithErr(code.ErrUserNotFound, c)
		return
	}

	if user.Enable != 1 {
		response.FailWithErr(code.ErrUserLocked, c)
		return
	}

	hasPasskey := user.Passkey != ""
	if strings.Contains(c.Request.Host, "127.0.0.1") || strings.Contains(c.Request.Host, "localhost") {
		hasPasskey = user.TestPasskey != ""
	}

	// 构建安全状态响应
	securityState := systemRes.SecurityStateRes{
		Exists:      true,
		HasTotp:     user.GoogleAuthSecret != "",
		HasPasskey:  hasPasskey,
		IsLocked:    user.Enable != 1,
		DisplayName: user.NickName,
	}

	response.OkWithDetailed(securityState, "Success", c)
}

// PasswordVerify 账号密码预验证（设计规范 5.2.2）
// @Tags     Auth
// @Summary  账号密码预验证
// @Produce  application/json
// @Param    data  body      systemReq.PasswordVerifyReq                                 true  "用户名、密码和验证码"
// @Success  200   {object}  response.Response{data=systemRes.PasswordVerifyRes,msg=string}  "验证结果"
// @Router   /auth/password/verify [post]
func (b *BaseApi) PasswordVerify(c *gin.Context) {
	var req systemReq.PasswordVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Username == "" || req.Password == "" {
		response.FailWithMessage("用户名和密码不能为空", c)
		return
	}

	// 验证图片验证码（防止爆破攻击）
	if req.Captcha == "" || req.CaptchaId == "" {
		response.FailWithMessage("请输入验证码", c)
		return
	}

	// 验证验证码
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}

	// 验证用户名和密码
	u := &system.SysUser{Username: req.Username, Password: req.Password}
	user, err := userService.Login(u)
	if err != nil {
		response.FailWithErr(code.ErrPasswordError, c)
		return
	}

	if user.Enable != 1 {
		response.FailWithErr(code.ErrUserLocked, c)
		return
	}

	if strings.Contains(c.Request.Host, "127.0.0.1") || strings.Contains(c.Request.Host, "localhost") {
		if user.GoogleAuthSecret != "" || user.TestPasskey != "" {
			response.FailWithMessage("该用户已绑定安全验证", c)
			return
		}
	} else {
		if user.GoogleAuthSecret != "" || user.Passkey != "" {
			response.FailWithMessage("该用户已绑定安全验证", c)
			return
		}
	}

	// 检查是否已绑定安全验证

	// 密码验证成功，生成绑定会话token
	sessionToken := utils.MD5V([]byte(fmt.Sprintf("%s_%d_%d", user.Username, user.ID, time.Now().Unix())))

	// 将用户信息存储到数据库中
	err = userService.CreateBindSession(sessionToken, user)
	if err != nil {
		global.HAB_LOG.Error("Failed to store bind session", zap.Error(err))
		response.FailWithMessage("系统错误", c)
		return
	}

	// 密码验证成功，可以进入绑定流程
	response.OkWithDetailed(systemRes.PasswordVerifyRes{
		Ok:          true,
		Next:        "bind_modal",
		BindSession: sessionToken,
	}, "验证成功", c)
}

func (b *BaseApi) PasswordLogin(c *gin.Context) {
	var req systemReq.PasswordVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Username == "" || req.Password == "" {
		response.FailWithErr(code.ErrUsernamePasswordEmpty, c)
		return
	}

	// 获取登录模式
	loginMode := global.HAB_CONFIG.System.LoginMode
	if loginMode == "" {
		loginMode = "simple"
	}

	// 验证码校验 — 仅 captcha 和 strict 模式需要
	if loginMode != "simple" {
		if req.Captcha == "" || req.CaptchaId == "" {
			response.FailWithErr(code.ErrCaptchaRequired, c)
			return
		}
		if !store.Verify(req.CaptchaId, req.Captcha, true) {
			response.FailWithErr(code.ErrCaptchaInvalid, c)
			return
		}
	}

	// 验证用户名和密码
	u := &system.SysUser{Username: req.Username, Password: req.Password}
	user, err := userService.Login(u)
	if err != nil {
		response.FailWithErr(code.ErrPasswordError, c)
		return
	}

	if user.Enable != 1 {
		response.FailWithErr(code.ErrUserLocked, c)
		return
	}

	// 角色检查 — 仅 strict 模式需要检查"账户状态"角色
	if loginMode == "strict" {
		var ok bool
		for _, authority := range user.Authorities {
			if authority.AuthorityName == "账户状态" {
				ok = true
				break
			}
		}
		if !ok {
			response.FailWithErr(code.ErrPermissionDenied, c)
			return
		}
	}

	// Token 签发 — strict 模式保持原行为，其他模式用 TokenNext 签发标准 JWT
	if loginMode == "strict" {
		token := pwd.SetPwdToken(user.ID)
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		}, "Login successful", c)
	} else {
		b.TokenNext(c, *user)
	}
}

// TotpLogin 用户名+TOTP登录（设计规范 5.2.3）
// @Tags     Auth
// @Summary  TOTP登录
// @Produce  application/json
// @Param    data  body      systemReq.TotpLoginReq                                      true  "用户名和TOTP验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "登录成功"
// @Router   /auth/totp/login [post]
func (b *BaseApi) TotpLogin(c *gin.Context) {
	var req systemReq.TotpLoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Username == "" || req.Code == "" {
		response.FailWithMessage("用户名和验证码不能为空", c)
		return
	}

	user, err := userService.GetUserByUsername(req.Username)
	if err != nil {
		response.FailWithErr(code.ErrUserNotFound, c)
		return
	}

	if user.Enable != 1 {
		response.FailWithErr(code.ErrUserLocked, c)
		return
	}

	// 检查是否绑定了TOTP
	if user.GoogleAuthSecret == "" {
		response.FailWithErr(code.ErrTotpNotBound, c)
		return
	}

	// 验证TOTP验证码
	valid := totp.Validate(req.Code, user.GoogleAuthSecret)
	if !valid {
		response.FailWithErr(code.ErrTotpCodeError, c)
		return
	}

	// 登录成功，生成token
	b.TokenNext(c, *user)
}

// PasskeyAssertionOptions 发起Passkey登录（设计规范 5.2.4）
// @Tags     Auth
// @Summary  获取Passkey登录选项
// @Produce  application/json
// @Param    data  body      systemReq.PasskeyAssertionOptionsReq                        true  "用户名（可选）"
// @Success  200   {object}  response.Response{data=systemRes.PasskeyAssertionOptionsRes,msg=string}  "Passkey登录选项"
// @Router   /auth/passkey/assertion/options [post]
func (b *BaseApi) PasskeyAssertionOptions(c *gin.Context) {
	var req systemReq.PasskeyAssertionOptionsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rpid := c.Request.Host
	if strings.Contains(rpid, "127.0.0.1") {
		rpid = "localhost"
	}

	// 生成Passkey登录选项
	publicKey := systemRes.PasskeyPublicKeyCredentialRequestOptions{
		Challenge:        generateChallenge(),
		RpId:             rpid,
		Timeout:          60000,
		UserVerification: "preferred",
	}

	// 如果提供了用户名，获取该用户的凭证
	if req.Username != "" {
		user, err := userService.GetUserByUsername(req.Username)
		if err == nil && user.Passkey != "" {
			publicKey.AllowCredentials = []systemRes.PasskeyPublicKeyCredentialDescriptor{{
				Id:         user.Passkey, // 这里应该是实际的credential ID
				Type:       "public-key",
				Transports: []string{"internal"},
			}}
		}
	}

	response.OkWithDetailed(systemRes.PasskeyAssertionOptionsRes{
		PublicKey: publicKey,
	}, "Success", c)
}

// PasskeyAssertionVerify 完成Passkey登录（设计规范 5.2.5）
// @Tags     Auth
// @Summary  验证Passkey登录
// @Produce  application/json
// @Param    data  body      systemReq.PasskeyAssertionVerifyReq                         true  "Passkey验证数据"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "登录成功"
// @Router   /auth/passkey/assertion/verify [post]
func (b *BaseApi) PasskeyAssertionVerify(c *gin.Context) {
	var req systemReq.PasskeyAssertionVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Id == "" {
		response.FailWithMessage("凭证ID不能为空", c)
		return
	}

	isLocal := strings.Contains(c.Request.Host, "127.0.0.1") || strings.Contains(c.Request.Host, "localhost")

	// 根据凭证ID查找对应的用户
	user, err := userService.GetUserByPasskeyCredentialId(req.Id, isLocal)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithErr(code.ErrPasskeyNotBound, c)
		} else {
			global.HAB_LOG.Error("Failed to get user by passkey credential", zap.String("credentialId", req.Id), zap.Error(err))
			response.FailWithMessage("系统错误", c)
		}
		return
	}

	if user.Enable != 1 {
		response.FailWithErr(code.ErrUserLocked, c)
		return
	}

	if user.GetPasskey(isLocal) == "" {
		response.FailWithErr(code.ErrPasskeyNotBound, c)
		return
	}

	// 实际项目中需要验证签名
	global.HAB_LOG.Info("Passkey login successful", zap.String("user", user.Username))

	// 登录成功，生成token
	b.TokenNext(c, *user)
}

// TotpBindInit 绑定TOTP初始化（设计规范 5.2.6）
// @Tags     Auth
// @Summary  初始化TOTP绑定
// @Produce  application/json
// @Success  200   {object}  response.Response{data=systemRes.TotpBindInitRes,msg=string}  "TOTP绑定信息"
// @Router   /auth/totp/bind/init [post]
func (b *BaseApi) TotpBindInit(c *gin.Context) {
	// 这里应该验证用户已通过密码验证
	// 简化处理，实际应该从session或临时token中获取用户信息

	// 生成密钥
	secret := base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%s", time.Now().UnixNano(), generateChallenge())))
	secret = secret[:32]

	// 生成otpauth URL
	issuer := global.HAB_CONFIG.AutoCode.Module
	username := "user" // 实际应该从上下文获取
	otpauthUrl := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s&digits=6&period=30",
		issuer, username, secret, issuer)

	response.OkWithDetailed(systemRes.TotpBindInitRes{
		QrcodeUrl:    otpauthUrl, // 前端用此生成二维码
		OtpauthUrl:   otpauthUrl,
		SecretMasked: secret[:8] + "****",
		Secret:       secret, // 临时返回，实际应该存在session中
	}, "Success", c)
}

// TotpBindVerify 验证TOTP绑定（设计规范 5.2.6）
// @Tags     Auth
// @Summary  验证TOTP绑定
// @Produce  application/json
// @Param    data  body      systemReq.TotpBindVerifyReq                                 true  "TOTP验证码"
// @Success  200   {object}  response.Response{data=systemRes.TotpBindVerifyRes,msg=string}  "绑定结果"
// @Router   /auth/totp/bind/verify [post]
func (b *BaseApi) TotpBindVerify(c *gin.Context) {
	var req systemReq.TotpBindVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Code == "" {
		response.FailWithMessage("验证码不能为空", c)
		return
	}

	// 从数据库获取绑定会话
	session, err := userService.GetBindSession(req.BindSession)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("绑定会话不存在或已失效", c)
		} else {
			global.HAB_LOG.Error("Failed to get bind session", zap.Error(err))
			response.FailWithMessage("系统错误", c)
		}
		return
	}

	// 验证TOTP验证码
	valid := totp.Validate(req.Code, req.Secret)
	if !valid {
		response.FailWithErr(code.ErrTotpCodeError, c)
		return
	}

	// 保存Passkey凭证到数据库
	err = userService.BindGoogleAuth(session.UserID, req.Secret)
	if err != nil {
		global.HAB_LOG.Error("Failed to save passkey", zap.Error(err))
		response.FailWithMessage("保存Passkey失败", c)
		return
	}

	// 清除绑定会话
	err = userService.DeleteBindSession(req.BindSession)
	if err != nil {
		global.HAB_LOG.Error("Failed to delete bind session", zap.Error(err))
		// 不阻止绑定流程，只记录错误
	}
	// 绑定成功，更新用户记录
	// 这里应该更新数据库中的用户记录
	response.OkWithDetailed(systemRes.TotpBindVerifyRes{
		Bound: true,
	}, "绑定成功", c)
}

// PasskeyAttestationOptions 绑定Passkey初始化（设计规范 5.2.7）
// @Tags     Auth
// @Summary  获取Passkey绑定选项
// @Produce  application/json
// @Param    data  body      systemReq.PasskeyAttestationOptionsReq                      true  "显示名称"
// @Success  200   {object}  response.Response{data=systemRes.PasskeyAttestationOptionsRes,msg=string}  "Passkey绑定选项"
// @Router   /auth/passkey/attestation/options [post]
func (b *BaseApi) PasskeyAttestationOptions(c *gin.Context) {
	var req systemReq.PasskeyAttestationOptionsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rpid := c.Request.Host
	if strings.Contains(rpid, "127.0.0.1") {
		rpid = "localhost"
	}
	var publicKey systemRes.PasskeyPublicKeyCredentialCreationOptions
	if req.DisplayName == "" {
		publicKey = systemRes.PasskeyPublicKeyCredentialCreationOptions{
			Challenge: generateChallenge(),
			//Rp: systemRes.PasskeyRelyingParty{
			//	Name: enum.AppKey,
			//	Id:   rpid,
			//},
			User: systemRes.PasskeyUser{
				Id: "dummy",
			},
			PubKeyCredParams: []systemRes.PasskeyPubKeyCredParam{
				{Type: "public-key", Alg: -7},   // ES256
				{Type: "public-key", Alg: -257}, // RS256
			},
			//Timeout:     60000,
			//Attestation: "none",
			//AuthenticatorSelection: systemRes.PasskeyAuthenticatorSelection{
			//	ResidentKey: "preferred",
			//	//UserVerification: "preferred",
			//},
		}
	} else {
		publicKey = systemRes.PasskeyPublicKeyCredentialCreationOptions{
			Challenge: generateChallenge(),
			Rp: systemRes.PasskeyRelyingParty{
				Name: global.HAB_CONFIG.AutoCode.Module,
				Id:   rpid,
			},
			User: systemRes.PasskeyUser{
				//Id: "dummy",
				Id:          base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s-%s", req.DisplayName, enum.AppKey, global.HAB_CONFIG.System.Environment))),
				Name:        fmt.Sprintf("%s:%s-%s", req.DisplayName, enum.AppKey, global.HAB_CONFIG.System.Environment),
				DisplayName: req.DisplayName,
			},
			PubKeyCredParams: []systemRes.PasskeyPubKeyCredParam{
				{Type: "public-key", Alg: -7},   // ES256
				{Type: "public-key", Alg: -257}, // RS256
			},
			Timeout:     60000,
			Attestation: "none",
			AuthenticatorSelection: systemRes.PasskeyAuthenticatorSelection{
				ResidentKey:      "preferred",
				UserVerification: "preferred",
			},
		}
	}

	response.OkWithDetailed(systemRes.PasskeyAttestationOptionsRes{
		PublicKey: publicKey,
	}, "Success", c)
}

// PasskeyAttestationVerify 验证Passkey绑定（设计规范 5.2.7）
// @Tags     Auth
// @Summary  验证Passkey绑定
// @Produce  application/json
// @Param    data  body      systemReq.PasskeyAttestationVerifyReq                       true  "Passkey绑定数据"
// @Success  200   {object}  response.Response{data=systemRes.PasskeyAttestationVerifyRes,msg=string}  "绑定结果"
// @Router   /auth/passkey/attestation/verify [post]
func (b *BaseApi) PasskeyAttestationVerify(c *gin.Context) {
	var req systemReq.PasskeyAttestationVerifyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Id == "" {
		response.FailWithMessage("凭证ID不能为空", c)
		return
	}

	if req.BindSession == "" {
		response.FailWithMessage("绑定会话无效", c)
		return
	}

	// 从数据库获取绑定会话
	session, err := userService.GetBindSession(req.BindSession)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("绑定会话不存在或已失效", c)
		} else {
			global.HAB_LOG.Error("Failed to get bind session", zap.Error(err))
			response.FailWithMessage("系统错误", c)
		}
		return
	}

	// 简化的Passkey绑定验证 - 实际应该解析attestationObject并验证
	global.HAB_LOG.Info("Passkey binding successful",
		zap.Uint("userId", session.UserID),
		zap.String("username", session.Username),
		zap.String("credentialId", req.Id))

	// 保存Passkey凭证到数据库
	err = userService.BindPasskey(session.UserID, req.Id, strings.Contains(c.Request.Host, "127.0.0.1") || strings.Contains(c.Request.Host, "localhost"))
	if err != nil {
		global.HAB_LOG.Error("Failed to save passkey", zap.Error(err))
		response.FailWithMessage("保存Passkey失败", c)
		return
	}

	// 清除绑定会话
	err = userService.DeleteBindSession(req.BindSession)
	if err != nil {
		global.HAB_LOG.Error("Failed to delete bind session", zap.Error(err))
		// 不阻止绑定流程，只记录错误
	}

	response.OkWithDetailed(systemRes.PasskeyAttestationVerifyRes{
		Bound: true,
	}, "绑定成功", c)
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	token, claims, err := utils.LoginToken(&user)
	if err != nil {
		global.HAB_LOG.Error("Failed to get token!", zap.Error(err))
		response.FailWithMessage("Failed to get token", c)
		return
	}
	if !global.HAB_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.HAB_LOG.Error("Failed to set login status!", zap.Error(err))
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	} else if err != nil {
		global.HAB_LOG.Error("Failed to set login status!", zap.Error(err))
		response.FailWithMessage("Failed to set login status", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("Failed to invalidate JWT", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	}
}

// 保留原有的用户管理功能

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	tep, pas := utils.GetUserTypeInfo(c)
	if tep != enum.SysUserTypeAdmin {
		r.Type, r.Parameter = tep, pas
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email, Type: r.Type, Parameter: r.Parameter}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.HAB_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithErr(code.ErrUserExists, c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{HAB_MODEL: global.HAB_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.HAB_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithErr(code.ErrOldPasswordError, c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := userService.GetUserInfoList(pageInfo, c)
	if err != nil {
		global.HAB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithErr(code.ErrGetInfoFailed, c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Success", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.HAB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithErr(code.ErrGetInfoFailed, c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "Success", c)
}

// BindGoogleAuthByLogin 登录前绑定谷歌验证器（保留兼容性）
// @Tags      SysUser
// @Summary   登录前绑定谷歌验证器
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindGoogleAuthByLogin true  "用户名, 密码, 验证码, 密钥"
// @Success   200   {object}  response.Response{msg=string}        "绑定谷歌验证器"
// @Router    /user/bindGoogleAuthByLogin [post]
func (b *BaseApi) BindGoogleAuthByLogin(c *gin.Context) {
	var req systemReq.BindGoogleAuthByLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证用户名和密码
	u := &system.SysUser{Username: req.Username, Password: req.Password}
	user, err := userService.Login(u)
	if err != nil {
		response.FailWithErr(code.ErrUserNotFound, c)
		return
	}

	// 检查是否已经绑定
	if user.GoogleAuthSecret != "" {
		response.FailWithMessage("该用户已绑定谷歌验证器，无法重复绑定", c)
		return
	}

	// 验证谷歌验证码
	valid := totp.Validate(req.Code, req.SecretKey)
	if !valid {
		response.FailWithErr(code.ErrTotpCodeError, c)
		return
	}

	// 绑定谷歌验证器
	err = userService.BindGoogleAuth(user.ID, req.SecretKey)
	if err != nil {
		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
		return
	}

	response.OkWithMessage("绑定成功", c)
}

// BindPasskeyByLogin 登录前绑定Passkey（保留兼容性）
// @Tags      SysUser
// @Summary   登录前绑定Passkey
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindPasskeyByLogin true  "用户名, 密码, Passkey数据"
// @Success   200   {object}  response.Response{msg=string}     "绑定Passkey"
// @Router    /user/bindPasskeyByLogin [post]
func (b *BaseApi) BindPasskeyByLogin(c *gin.Context) {
	var req systemReq.BindPasskeyByLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证用户名和密码
	u := &system.SysUser{Username: req.Username, Password: req.Password}
	user, err := userService.Login(u)
	if err != nil {
		response.FailWithErr(code.ErrUserNotFound, c)
		return
	}

	// 检查是否已经绑定
	if user.Passkey != "" {
		response.FailWithMessage("该用户已绑定通行密钥，无法重复绑定", c)
		return
	}

	// 绑定Passkey
	err = userService.BindPasskey(user.ID, req.PasskeyData, strings.Contains(c.Request.Host, "127.0.0.1") || strings.Contains(c.Request.Host, "localhost"))
	if err != nil {
		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
		return
	}

	response.OkWithMessage("绑定成功", c)
}

// Login 保留原接口以兼容现有代码
// @Tags     Base
// @Summary  用户登录（兼容接口）
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 兼容原有登录逻辑，引导用户使用新的API
	response.FailWithMessage("请使用新的登录接口", c)
}

//// 以下是保持兼容性的旧API方法 - 简化实现
//
//func (b *BaseApi) SetUserAuthority(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) DeleteUser(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) SetUserInfo(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) SetSelfInfo(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) ResetPassword(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) SetSelfSetting(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) GetGoogleAuthQR(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) BindGoogleAuth(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) ResetGoogleAuth(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) BindPasskey(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
//
//func (b *BaseApi) UnbindSecurity(c *gin.Context) {
//	response.FailWithMessage("此功能暂时不可用，请联系管理员", c)
//}
