package system

import (
	"hab/enum"
	"hab/global"
	"hab/model/common"

	"github.com/google/uuid"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetType() enum.SysUserType
	GetParameter() string
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	global.HAB_MODEL
	UUID          uuid.UUID        `json:"uuid" gorm:"index;comment:用户UUID"`                                                                   // 用户UUID
	Username      string           `json:"userName" gorm:"index;comment:用户登录名"`                                                                // 用户登录名
	Password      string           `json:"-"  gorm:"comment:用户登录密码"`                                                                           // 用户登录密码
	NickName      string           `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                          // 用户昵称
	HeaderImg     string           `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/HAB_header.jpg;comment:用户头像"`               // 用户头像
	Authority     SysAuthority     `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`                        // 用户角色
	AuthorityId   uint             `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                                      // 用户角色ID
	SideMode      string           `json:"sideMode" gorm:"default:dark;comment:用户角色ID"`                                                        // 用户侧边主题
	ActiveColor   string           `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"`                                                  // 活跃颜色
	BaseColor     string           `json:"baseColor" gorm:"default:#fff;comment:用户角色ID"`                                                       // 基础颜色
	Authorities   []SysAuthority   `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                   // 多用户角色
	Phone         string           `json:"phone"  gorm:"comment:用户手机号"`                                                                        // 用户手机号
	Email         string           `json:"email"  gorm:"comment:用户邮箱"`                                                                         // 用户邮箱
	Enable        int              `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                                    //用户是否被冻结 1正常 2冻结
	OriginSetting common.JSONMap   `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	Type          enum.SysUserType `json:"type" gorm:"default:0;comment:用户类型 1普通用户 2商户"`                                                       //用户类型 1普通用户 2商户
	Parameter     string           `json:"parameter" gorm:"type:text;comment:用户参数"`                                                            //用户参数
	Language      string           `json:"language" gorm:"default:zh-CN;comment:用户语言"`                                                         // 用户语言
	//GoogleAuthEnabled bool             `json:"googleAuthEnabled" gorm:"default:false;comment:是否启用谷歌验证器"`                                           // 是否启用谷歌验证器
	GoogleAuthSecret string `json:"-" gorm:"comment:谷歌验证器密钥"`                         // 谷歌验证器密钥
	Passkey          string `json:"passkey" gorm:"comment:用户Passkey;size:2000"`       // 用户Passkey
	TestPasskey      string `json:"testPasskey" gorm:"comment:用户测试Passkey;size:2000"` // 用户测试Passkey
}

func (s *SysUser) GetPasskey(isLocal bool) string {
	if isLocal {
		return s.TestPasskey
	}
	return s.Passkey
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}

func (s *SysUser) GetType() enum.SysUserType {
	return s.Type
}

func (s *SysUser) GetParameter() string {
	return s.Parameter
}

// SysBindSession 用户绑定会话表
type SysBindSession struct {
	global.HAB_MODEL
	SessionToken string `json:"session_token" gorm:"uniqueIndex;comment:会话token"`
	UserID       uint   `json:"user_id" gorm:"comment:用户ID"`
	Username     string `json:"username" gorm:"comment:用户名"`
	NickName     string `json:"nick_name" gorm:"comment:用户昵称"`
}

func (SysBindSession) TableName() string {
	return "sys_bind_sessions"
}
