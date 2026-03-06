package enum

type SysUserType int

const (
	SysUserTypeAdmin    SysUserType = 0 // 管理员
	SysUserTypeAgent    SysUserType = 1 // Agent
	SysUserTypeMerchant SysUserType = 2 // 商户
)
