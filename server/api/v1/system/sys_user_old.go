package system

import (
	"strconv"
	"time"

	"hab/code"
	"hab/global"
	"hab/model/common"
	"hab/model/common/request"
	"hab/model/common/response"
	"hab/model/system"
	systemReq "hab/model/system/request"
	systemRes "hab/model/system/response"
	"hab/utils"

	"encoding/base32"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 移除这些结构体，将使用 systemRes 包中的定义

// generateChallenge 生成WebAuthn挑战
//func generateChallenge() string {
//	challenge := make([]byte, 32)
//	_, err := rand.Read(challenge)
//	if err != nil {
//		// 如果随机数生成失败，使用时间戳作为后备
//		timeBytes := []byte(fmt.Sprintf("%d", time.Now().UnixNano()))
//		if len(timeBytes) < 32 {
//			paddedBytes := make([]byte, 32)
//			copy(paddedBytes, timeBytes)
//			challenge = paddedBytes
//		} else {
//			challenge = timeBytes[:32]
//		}
//	}
//	return base64.RawURLEncoding.EncodeToString(challenge)
//}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
//func (b *BaseApi) Login(c *gin.Context) {
//	var l systemReq.Login
//	err := c.ShouldBindJSON(&l)
//	//key := c.ClientIP()
//
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	err = utils.Verify(l, utils.LoginVerify)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	// 根据登录类型决定是否需要验证密码
//	var user *system.SysUser
//	if l.LoginType == "Passkey" || (l.LoginType == "AuthCode" && l.GoogleAuthCode != "") {
//		// Passkey登录或提供谷歌验证码时不需要密码验证，只验证用户名
//		user, err = userService.GetUserByUsername(l.Username)
//		if err != nil {
//			response.FailWithErr(code.ErrUserNotFound, c)
//			return
//		}
//	} else {
//		// 其他情况需要密码验证
//		u := &system.SysUser{Username: l.Username, Password: l.Password}
//		user, err = userService.Login(u)
//		if err != nil {
//			response.FailWithErr(code.ErrUserNotFound, c)
//			return
//		}
//	}
//	if user.Enable != 1 {
//		global.HAB_LOG.Error("Login failed! User is disabled!")
//		response.FailWithErr(code.ErrUserDisabled, c)
//		return
//	}
//
//	if user.GoogleAuthSecret == "" && user.Passkey == "" {
//		// 生成密钥
//		secret := base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%s", user.ID, time.Now().String())))
//		secret = secret[:32]
//
//		// 生成二维码内容
//		issuer := global.HAB_CONFIG.AutoCode.Module
//		qrCode := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, user.Username, secret, issuer)
//		response.FailWithErrData(code.ErrUnsafe, BindKeyInfo{
//			PassKeyInfo: PassKeyInfo{
//				Challenge: generateChallenge(),
//				RpId:      l.Domain,
//				RpName:    enum.AppKey,
//				User: struct {
//					Id          string `json:"id"`
//					Name        string `json:"name"`
//					DisplayName string `json:"displayName"`
//				}{
//					Id:          strconv.Itoa(int(user.ID)),
//					Name:        fmt.Sprintf("%s:%s-%s", user.Username, enum.AppKey, global.HAB_CONFIG.System.Environment),
//					DisplayName: fmt.Sprintf("%s:%s-%s", user.NickName, enum.AppKey, global.HAB_CONFIG.System.Environment),
//				},
//			},
//			GoogleAuthInfo: GoogleAuthInfo{
//				QRCode:    qrCode,
//				SecretKey: secret,
//			},
//		}, c)
//		return
//	}
//
//	switch l.LoginType {
//	case "AuthCode":
//		//验证验证码
//		if l.GoogleAuthCode == "" {
//			err = code.ErrGoogleAuthRequired
//		} else {
//			valid := totp.Validate(l.GoogleAuthCode, user.GoogleAuthSecret)
//			if !valid {
//				err = code.ErrInvalidGoogleAuth
//			}
//		}
//	case "Passkey":
//		// 验证Passkey
//		if l.Passkey == "" {
//			// 如果没有提供Passkey数据，返回挑战信息
//			response.FailWithErrData(code.ErrPasskeyRequired, BindKeyInfo{
//				PassKeyInfo: PassKeyInfo{
//					Challenge: generateChallenge(),
//					RpId:      l.Domain,
//					RpName:    enum.AppKey,
//					User: struct {
//						Id          string `json:"id"`
//						Name        string `json:"name"`
//						DisplayName string `json:"displayName"`
//					}{
//						Id:          strconv.Itoa(int(user.ID)),
//						Name:        user.Username,
//						DisplayName: fmt.Sprintf("%s:%s-%s", user.NickName, enum.AppKey, global.HAB_CONFIG.System.Environment),
//					},
//				},
//			}, c)
//			return
//		} else {
//			// 简单验证Passkey数据（实际应用中需要更复杂的验证）
//			if user.Passkey == "" {
//				err = code.ErrPasskeyRequired
//			} else {
//				// 这里应该验证Passkey的签名等，但为了简化，我们暂时只检查用户是否已绑定
//				global.HAB_LOG.Info("Passkey login attempt", zap.String("user", user.Username))
//			}
//		}
//	default:
//		// 默认需要验证码
//		if user.GoogleAuthSecret != "" {
//			if l.GoogleAuthCode == "" {
//				err = code.ErrGoogleAuthRequired
//			} else {
//				valid := totp.Validate(l.GoogleAuthCode, user.GoogleAuthSecret)
//				if !valid {
//					err = code.ErrInvalidGoogleAuth
//				}
//			}
//		} else if user.Passkey != "" {
//			err = code.ErrPasskeyRequired
//		} else {
//			err = code.ErrUnsafe
//		}
//	}
//	if err != nil {
//		global.HAB_LOG.Error("Login failed!", zap.Error(err))
//		response.FailWithErr(err, c)
//		return
//	}
//	b.TokenNext(c, *user)
//	return
//}

// TokenNext 登录以后签发jwt
//func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
//	token, claims, err := utils.LoginToken(&user)
//	if err != nil {
//		global.HAB_LOG.Error("Failed to get token!", zap.Error(err))
//		response.FailWithMessage("Failed to get token", c)
//		return
//	}
//	if !global.HAB_CONFIG.System.UseMultipoint {
//		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
//		response.OkWithDetailed(systemRes.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
//		}, "Login successful", c)
//		return
//	}
//
//	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
//		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
//			global.HAB_LOG.Error("Failed to set login status!", zap.Error(err))
//			response.FailWithMessage("Failed to set login status", c)
//			return
//		}
//		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
//		response.OkWithDetailed(systemRes.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
//		}, "Login successful", c)
//	} else if err != nil {
//		global.HAB_LOG.Error("Failed to set login status!", zap.Error(err))
//		response.FailWithMessage("Failed to set login status", c)
//	} else {
//		var blackJWT system.JwtBlacklist
//		blackJWT.Jwt = jwtStr
//		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
//			response.FailWithMessage("Failed to invalidate JWT", c)
//			return
//		}
//		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
//			response.FailWithMessage("Failed to set login status", c)
//			return
//		}
//		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
//		response.OkWithDetailed(systemRes.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
//		}, "Login successful", c)
//	}
//}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
//func (b *BaseApi) Register(c *gin.Context) {
//	var r systemReq.Register
//	err := c.ShouldBindJSON(&r)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	err = utils.Verify(r, utils.RegisterVerify)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	var authorities []system.SysAuthority
//	for _, v := range r.AuthorityIds {
//		authorities = append(authorities, system.SysAuthority{
//			AuthorityId: v,
//		})
//	}
//	tep, pas := utils.GetUserTypeInfo(c)
//	if tep != enum.SysUserTypeAdmin {
//		r.Type, r.Parameter = tep, pas
//	}
//	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email, Type: r.Type, Parameter: r.Parameter}
//	userReturn, err := userService.Register(*user)
//	if err != nil {
//		global.HAB_LOG.Error("注册失败!", zap.Error(err))
//		response.FailWithErr(code.ErrUserExists, c)
//		return
//	}
//	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
//}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
//func (b *BaseApi) ChangePassword(c *gin.Context) {
//	var req systemReq.ChangePasswordReq
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	err = utils.Verify(req, utils.ChangePasswordVerify)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	uid := utils.GetUserID(c)
//	u := &system.SysUser{HAB_MODEL: global.HAB_MODEL{ID: uid}, Password: req.Password}
//	_, err = userService.ChangePassword(u, req.NewPassword)
//	if err != nil {
//		global.HAB_LOG.Error("修改失败!", zap.Error(err))
//		response.FailWithErr(code.ErrOldPasswordError, c)
//		return
//	}
//	response.OkWithMessage("修改成功", c)
//}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
//func (b *BaseApi) GetUserList(c *gin.Context) {
//	var pageInfo systemReq.GetUserList
//	err := c.ShouldBindJSON(&pageInfo)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	err = utils.Verify(pageInfo, utils.PageInfoVerify)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	list, total, err := userService.GetUserInfoList(pageInfo, c)
//	if err != nil {
//		global.HAB_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithErr(code.ErrGetInfoFailed, c)
//		return
//	}
//	response.OkWithDetailed(response.PageResult{
//		List:     list,
//		Total:    total,
//		Page:     pageInfo.Page,
//		PageSize: pageInfo.PageSize,
//	}, "Success", c)
//}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.HAB_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithErr(code.ErrSetInfoFailed, c)
		return
	}
	claims := utils.GetUserInfo(c)
	claims.AuthorityId = sua.AuthorityId
	token, err := utils.NewJWT().CreateToken(*claims)
	if err != nil {
		global.HAB_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Header("new-token", token)
	c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
	utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	response.OkWithMessage("修改成功", c)
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authorityID := utils.GetUserAuthorityId(c)
	err = userService.SetUserAuthorities(authorityID, sua.ID, sua.AuthorityIds)
	if err != nil {
		global.HAB_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithErr(code.ErrDeleteSelfNotAllowed, c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.HAB_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Success", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(user.AuthorityIds) != 0 {
		authorityID := utils.GetUserAuthorityId(c)
		err = userService.SetUserAuthorities(authorityID, user.ID, user.AuthorityIds)
		if err != nil {
			global.HAB_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithErr(code.ErrSetInfoFailed, c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		HAB_MODEL: global.HAB_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
		Language:  user.Language,
		Type:      user.Type,
		Parameter: user.Parameter,
	})
	if err != nil {
		global.HAB_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithErr(code.ErrSetInfoFailed, c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		HAB_MODEL: global.HAB_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.HAB_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfSetting
// @Tags      SysUser
// @Summary   设置用户配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      map[string]interface{}  true  "用户配置数据"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户配置"
// @Router    /user/SetSelfSetting [put]
func (b *BaseApi) SetSelfSetting(c *gin.Context) {
	var req common.JSONMap
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetSelfSetting(req, utils.GetUserID(c))
	if err != nil {
		global.HAB_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
//func (b *BaseApi) GetUserInfo(c *gin.Context) {
//	uuid := utils.GetUserUuid(c)
//	ReqUser, err := userService.GetUserInfo(uuid)
//	if err != nil {
//		global.HAB_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithErr(code.ErrGetInfoFailed, c)
//		return
//	}
//	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "Success", c)
//}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.HAB_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithErr(code.ErrResetPasswordFailed, c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

// GetGoogleAuthQR
// @Tags      SysUser
// @Summary   获取谷歌验证器二维码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.GoogleAuthResponse,msg=string}  "获取谷歌验证器二维码"
// @Router    /user/getGoogleAuthQR [get]
func (b *BaseApi) GetGoogleAuthQR(c *gin.Context) {
	userID := utils.GetUserID(c)
	user, err := userService.GetUserByID(userID)
	if err != nil {
		global.HAB_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	// 生成密钥
	secret := base32.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%s", userID, time.Now().String())))
	secret = secret[:32]

	// 生成二维码内容
	issuer := global.HAB_CONFIG.AutoCode.Module
	qrCode := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, user.Username, secret, issuer)

	response.OkWithDetailed(systemRes.GoogleAuthResponse{
		QRCode:    qrCode,
		SecretKey: secret,
	}, "Success", c)
}

// BindGoogleAuth
// @Tags      SysUser
// @Summary   绑定谷歌验证器
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindGoogleAuth      true  "验证码, 密钥"
// @Success   200   {object}  response.Response{msg=string}  "绑定谷歌验证器"
// @Router    /user/bindGoogleAuth [post]
//func (b *BaseApi) BindGoogleAuth(c *gin.Context) {
//	var req systemReq.BindGoogleAuth
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	// 验证验证码
//	valid := totp.Validate(req.Code, req.SecretKey)
//	if !valid {
//		response.FailWithErr(code.ErrInvalidGoogleAuth, c)
//		return
//	}
//
//	userID := utils.GetUserID(c)
//	err = userService.BindGoogleAuth(userID, req.SecretKey)
//	if err != nil {
//		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
//		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
//		return
//	}
//	response.OkWithMessage("绑定成功", c)
//}

// ResetGoogleAuth
// @Tags      SysUser
// @Summary   重置谷歌验证器
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "重置谷歌验证器"
// @Router    /user/resetGoogleAuth [post]
func (b *BaseApi) ResetGoogleAuth(c *gin.Context) {
	userID := utils.GetUserID(c)
	err := userService.ResetGoogleAuth(userID)
	if err != nil {
		global.HAB_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithErr(code.ErrResetGoogleAuthFailed, c)
		return
	}

	response.OkWithMessage("重置成功", c)
}

// BindPasskey
// @Tags      SysUser
// @Summary   绑定Passkey
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindPasskey         true  "Passkey数据"
// @Success   200   {object}  response.Response{msg=string}  "绑定Passkey"
// @Router    /user/bindPasskey [post]
//func (b *BaseApi) BindPasskey(c *gin.Context) {
//	var req systemReq.BindPasskey
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	userID := utils.GetUserID(c)
//	err = userService.BindPasskey(userID, req.PasskeyData)
//	if err != nil {
//		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
//		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
//		return
//	}
//	response.OkWithMessage("绑定成功", c)
//}

// UnbindSecurity
// @Tags      SysUser
// @Summary   解绑安全验证
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById               true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "解绑安全验证"
// @Router    /user/unbindSecurity [post]
func (b *BaseApi) UnbindSecurity(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.UnbindSecurity(uint(reqId.ID))
	if err != nil {
		global.HAB_LOG.Error("解绑失败!", zap.Error(err))
		response.FailWithMessage("解绑失败", c)
		return
	}
	response.OkWithMessage("解绑成功", c)
}

// BindGoogleAuthByLogin
// @Tags      SysUser
// @Summary   登录前绑定谷歌验证器
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindGoogleAuthByLogin true  "用户名, 密码, 验证码, 密钥"
// @Success   200   {object}  response.Response{msg=string}        "绑定谷歌验证器"
// @Router    /user/bindGoogleAuthByLogin [post]
//func (b *BaseApi) BindGoogleAuthByLogin(c *gin.Context) {
//	var req systemReq.BindGoogleAuthByLogin
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	// 验证用户名和密码
//	u := &system.SysUser{Username: req.Username, Password: req.Password}
//	user, err := userService.Login(u)
//	if err != nil {
//		response.FailWithErr(code.ErrUserNotFound, c)
//		return
//	}
//
//	// 检查是否已经绑定
//	if user.GoogleAuthSecret != "" {
//		response.FailWithMessage("该用户已绑定谷歌验证器，无法重复绑定", c)
//		return
//	}
//
//	// 验证谷歌验证码
//	valid := totp.Validate(req.Code, req.SecretKey)
//	if !valid {
//		response.FailWithErr(code.ErrInvalidGoogleAuth, c)
//		return
//	}
//
//	// 绑定谷歌验证器
//	err = userService.BindGoogleAuth(user.ID, req.SecretKey)
//	if err != nil {
//		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
//		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
//		return
//	}
//
//	response.OkWithMessage("绑定成功", c)
//}

// BindPasskeyByLogin
// @Tags      SysUser
// @Summary   登录前绑定Passkey
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.BindPasskeyByLogin true  "用户名, 密码, Passkey数据"
// @Success   200   {object}  response.Response{msg=string}     "绑定Passkey"
// @Router    /user/bindPasskeyByLogin [post]
//func (b *BaseApi) BindPasskeyByLogin(c *gin.Context) {
//	var req systemReq.BindPasskeyByLogin
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	// 验证用户名和密码
//	u := &system.SysUser{Username: req.Username, Password: req.Password}
//	user, err := userService.Login(u)
//	if err != nil {
//		response.FailWithErr(code.ErrUserNotFound, c)
//		return
//	}
//
//	// 检查是否已经绑定
//	if user.Passkey != "" {
//		response.FailWithMessage("该用户已绑定通行密钥，无法重复绑定", c)
//		return
//	}
//
//	// 绑定Passkey
//	err = userService.BindPasskey(user.ID, req.PasskeyData)
//	if err != nil {
//		global.HAB_LOG.Error("绑定失败!", zap.Error(err))
//		response.FailWithErr(code.ErrBindGoogleAuthFailed, c)
//		return
//	}
//
//	response.OkWithMessage("绑定成功", c)
//}
