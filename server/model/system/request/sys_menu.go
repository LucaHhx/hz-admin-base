package request

import (
	"hab/global"
	"hab/model/system"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		HAB_MODEL: global.HAB_MODEL{ID: 53},
		ParentId:  0,
		Path:      "userInfo",
		Name:      "userInfo",
		Component: "view/person/person.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "用户信息",
			Icon:  "setting",
		},
	}}
}
