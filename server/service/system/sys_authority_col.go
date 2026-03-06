package system

import (
	"errors"
	"hab/global"
	"hab/model/system"
	"hab/model/system/request"
	"hab/model/system/response"

	"gorm.io/gorm"
)

type AuthorityColService struct{}

var AuthorityColServiceApp = new(AuthorityColService)

func (a *AuthorityColService) GetAuthorityCol(req request.SysAuthorityColReq) (res response.SysAuthorityColRes, err error) {
	var authorityCol []system.SysAuthorityCol
	err = global.HAB_DB.Find(&authorityCol, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityCol {
		selected = append(selected, v.SysTableColumnsID)
	}
	res.Selected = selected
	return res, err
}

func (a *AuthorityColService) SetAuthorityCol(req request.SysAuthorityColReq) (err error) {
	return global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		var authorityCol []system.SysAuthorityCol
		err = tx.Delete(&[]system.SysAuthorityCol{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityCol = append(authorityCol, system.SysAuthorityCol{
				AuthorityId:       req.AuthorityId,
				SysMenuID:         req.MenuID,
				SysTableColumnsID: v,
			})
		}
		if len(authorityCol) > 0 {
			err = tx.Create(&authorityCol).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (a *AuthorityColService) CanRemoveAuthorityCol(ID string) (err error) {
	fErr := global.HAB_DB.First(&system.SysAuthorityCol{}, "sys_table_columns_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("此列正在被使用无法删除")
}

// GetMenuColumns 获取菜单对应的所有列
func (a *AuthorityColService) GetMenuColumns(menuName string) (columns []system.SysTableColumns, err error) {
	err = global.HAB_DB.Where("struct_name = ?", menuName).Find(&columns).Error
	return columns, err
}
