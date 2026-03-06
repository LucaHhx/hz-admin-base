package response

// SysAuthorityColRes 角色列表权限返回数据结构体
type SysAuthorityColRes struct {
	Selected []uint `json:"selected"` // 选中的列ID
}
