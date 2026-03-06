package system

type SysAuthorityCol struct {
	AuthorityId       uint            `gorm:"comment:角色ID"`
	SysMenuID         uint            `gorm:"comment:菜单ID"`
	SysTableColumnsID uint            `gorm:"comment:列ID"`
	SysTableColumns   SysTableColumns `gorm:"comment:列详情"`
	//SysBaseMenuBtnID uint           `
}
