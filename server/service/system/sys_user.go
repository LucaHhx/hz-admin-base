package system

import (
	"errors"
	"fmt"
	"hab/enum"
	"hab/model/common"
	systemReq "hab/model/system/request"
	"time"

	"github.com/gin-gonic/gin"

	"hab/global"
	"hab/model/system"
	"hab/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.HAB_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.HAB_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.HAB_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.HAB_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.HAB_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.HAB_DB.Save(&user).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info systemReq.GetUserList, c *gin.Context) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.HAB_DB.Model(&system.SysUser{})
	var userList []system.SysUser

	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}

	tep, pas := utils.GetUserTypeInfo(c)
	if tep != enum.SysUserTypeAdmin {
		db = db.Where("type = ? and parameter = ? and id != ?", tep, pas, utils.GetUserID(c))
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {

	assignErr := global.HAB_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	var authority system.SysAuthority
	err = global.HAB_DB.Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return err
	}
	var authorityMenu []system.SysAuthorityMenu
	var authorityMenuIDs []string
	err = global.HAB_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
	if err != nil {
		return err
	}

	for i := range authorityMenu {
		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
	}

	var authorityMenus []system.SysBaseMenu
	err = global.HAB_DB.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
	if err != nil {
		return err
	}
	hasMenu := false
	for i := range authorityMenus {
		if authorityMenus[i].Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return errors.New("找不到默认路由,无法切换本角色")
	}

	err = global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(adminAuthorityID, id uint, authorityIds []uint) (err error) {
	return global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		TxErr := tx.Where("id = ?", id).First(&user).Error
		if TxErr != nil {
			global.HAB_LOG.Debug(TxErr.Error())
			return errors.New("查询用户数据失败")
		}
		TxErr = tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			e := AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityID, v)
			if e != nil {
				return e
			}
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Model(&user).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int) (err error) {
	return global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.HAB_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "enable", "language", "type", "parameter").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"enable":     req.Enable,
			"language":   req.Language,
			"type":       req.Type,
			"parameter":  req.Parameter,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.HAB_DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfSetting
//@description: 设置用户配置
//@param: req datatypes.JSON, uid uint
//@return: err error

func (userService *UserService) SetSelfSetting(req common.JSONMap, uid uint) error {
	return global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", uid).Update("origin_setting", req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.HAB_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.HAB_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.HAB_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ResetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

// GetUserByID 根据ID获取用户信息
func (userService *UserService) GetUserByID(id uint) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.HAB_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

// BindGoogleAuth 绑定谷歌验证器
func (userService *UserService) BindGoogleAuth(userID uint, secret string) error {
	return global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"google_auth_secret": secret,
	}).Error
}

// ResetGoogleAuth 重置谷歌验证器
func (userService *UserService) ResetGoogleAuth(userID uint) error {
	return global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", userID).Updates(map[string]interface{}{
		//"google_auth_enabled": false,
		"google_auth_secret": "",
	}).Error
}

// BindPasskey 绑定Passkey
func (userService *UserService) BindPasskey(userID uint, passkeyData string, isTest bool) error {
	if isTest {
		return global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", userID).Update("test_passkey", passkeyData).Error
	}
	return global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", userID).Update("passkey", passkeyData).Error
}

// GetUserByUsername 根据用户名获取用户信息（不验证密码）
func (userService *UserService) GetUserByUsername(username string) (*system.SysUser, error) {
	var user system.SysUser
	err := global.HAB_DB.Where("username = ?", username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err != nil {
		return nil, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&user)
	return &user, nil
}

// UnbindSecurity 解绑安全验证（谷歌验证器和Passkey）
func (userService *UserService) UnbindSecurity(userID uint) error {
	// 清理数据库中的安全绑定
	err := global.HAB_DB.Model(&system.SysUser{}).Where("id = ?", userID).Updates(map[string]interface{}{
		//"google_auth_enabled": false,
		"google_auth_secret": "",
		"passkey":            "",
		"test_passkey":       "",
	}).Error

	if err != nil {
		return err
	}

	// 清理数据库中该用户的所有绑定会话
	err = userService.DeleteUserBindSessions(userID)
	if err != nil {
		// 记录错误但不阻止解绑流程
		global.HAB_LOG.Error("Failed to delete user bind sessions", zap.Uint("userID", userID), zap.Error(err))
	}

	return nil
}

// CreateBindSession 创建绑定会话
func (userService *UserService) CreateBindSession(sessionToken string, user *system.SysUser) error {
	bindSession := system.SysBindSession{
		SessionToken: sessionToken,
		UserID:       user.ID,
		Username:     user.Username,
		NickName:     user.NickName,
	}
	return global.HAB_DB.Create(&bindSession).Error
}

// GetBindSession 获取绑定会话
func (userService *UserService) GetBindSession(sessionToken string) (*system.SysBindSession, error) {
	var session system.SysBindSession
	err := global.HAB_DB.Where("session_token = ?", sessionToken).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// DeleteBindSession 删除绑定会话
func (userService *UserService) DeleteBindSession(sessionToken string) error {
	return global.HAB_DB.Where("session_token = ?", sessionToken).Delete(&system.SysBindSession{}).Error
}

// DeleteUserBindSessions 删除用户的所有绑定会话
func (userService *UserService) DeleteUserBindSessions(userID uint) error {
	return global.HAB_DB.Where("user_id = ?", userID).Delete(&system.SysBindSession{}).Error
}

// GetUserByPasskeyCredentialId 根据Passkey凭证ID查找用户
func (userService *UserService) GetUserByPasskeyCredentialId(credentialId string, isTest bool) (*system.SysUser, error) {
	var user system.SysUser
	if isTest {
		err := global.HAB_DB.Where("test_passkey = ?", credentialId).Preload("Authorities").Preload("Authority").First(&user).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := global.HAB_DB.Where("passkey = ?", credentialId).Preload("Authorities").Preload("Authority").First(&user).Error
		if err != nil {
			return nil, err
		}
	}

	MenuServiceApp.UserAuthorityDefaultRouter(&user)
	return &user, nil
}
