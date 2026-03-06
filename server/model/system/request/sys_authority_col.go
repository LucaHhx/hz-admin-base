package request

// SysAuthorityColReq 角色列表权限请求参数结构体
type SysAuthorityColReq struct {
	AuthorityId uint   `json:"authorityId"` // 角色ID
	MenuID      uint   `json:"menuId"`      // 菜单ID
	Selected    []uint `json:"selected"`    // 选中的列ID
}
