package system

import (
	"hab/global"
	"hab/model/system"
	systemReq "hab/model/system/request"

	"gorm.io/gorm"
)

type SysTableColumnsService struct{}

// CreateSysTableColumns 创建表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) CreateSysTableColumns(sysTableColumns *system.SysTableColumns) (err error) {
	err = global.HAB_DB.Create(sysTableColumns).Error
	return err
}

// ImportSysTableColumns 导入表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) ImportSysTableColumns(sysTableColumnsList []system.SysTableColumns) (err error) {
	return global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		// 截断表（清空表并重置自增ID）
		if err := tx.Exec("TRUNCATE TABLE sys_table_columns").Error; err != nil {
			return err
		}
		// 批量写入
		if err := tx.Create(&sysTableColumnsList).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteSysTableColumns 删除表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) DeleteSysTableColumns(ID string, userID uint) (err error) {
	err = global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysTableColumns{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&system.SysTableColumns{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSysTableColumnsByIds 批量删除表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) DeleteSysTableColumnsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.HAB_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysTableColumns{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&system.SysTableColumns{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSysTableColumns 更新表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) UpdateSysTableColumns(sysTableColumns system.SysTableColumns) (err error) {
	err = global.HAB_DB.Model(&system.SysTableColumns{}).Where("id = ?", sysTableColumns.ID).Updates(&sysTableColumns).Error
	return err
}

// GetSysTableColumns 根据ID获取表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) GetSysTableColumns(ID string) (sysTableColumns system.SysTableColumns, err error) {
	err = global.HAB_DB.Where("id = ?", ID).First(&sysTableColumns).Error
	return
}

// GetSysTableColumnsInfoList 分页获取表字段配置记录
// Author [yourname](https://github.com/yourname)
func (sysTableColumnsService *SysTableColumnsService) GetSysTableColumnsInfoList(info systemReq.SysTableColumnsSearch) (list []system.SysTableColumns, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.HAB_DB.Model(&system.SysTableColumns{})
	var sysTableColumnss []system.SysTableColumns
	// If there is a condition search, the search statement will be automatically created below
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StructName != nil && *info.StructName != "" {
		db = db.Where("struct_name like ?", "%"+*info.StructName+"%")
	}
	if info.JsonName != nil && *info.JsonName != "" {
		db = db.Where("json_name = ?", *info.JsonName)
	}
	if info.ColumnName != nil && *info.ColumnName != "" {
		db = db.Where("column_name = ?", *info.ColumnName)
	}
	if info.StartWith != nil && info.EndWith != nil {
		db = db.Where("with BETWEEN ? AND ? ", info.StartWith, info.EndWith)
	}
	if info.Type != nil && *info.Type != "" {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Sortable != nil {
		db = db.Where("sortable = ?", *info.Sortable)
	}
	if info.Filter != nil {
		db = db.Where("filter = ?", *info.Filter)
	}
	if info.DefaultFilter != nil && *info.DefaultFilter != "" {
		db = db.Where("default_filter = ?", *info.DefaultFilter)
	}
	if info.FilterList != nil && *info.FilterList != "" {
		db = db.Where("filter_list = ?", *info.FilterList)
	}
	if info.MenuId != nil && *info.MenuId != 0 {
		db = db.Where("menu_id = ?", *info.MenuId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["struct_name"] = true
	orderMap["json_name"] = true
	orderMap["column_name"] = true
	orderMap["with"] = true
	orderMap["type"] = true
	orderMap["sortable"] = true
	orderMap["filter"] = true
	orderMap["default_filter"] = true
	orderMap["filter_list"] = true
	orderMap["menu_id"] = true
	orderMap["sort"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysTableColumnss).Error
	return sysTableColumnss, total, err
}
func (sysTableColumnsService *SysTableColumnsService) GetSysTableColumnsPublic() {
	// This method is for getting data source defined data
	// Please implement it yourself
}

func (sysTableColumnsService *SysTableColumnsService) GetStructNameColumns(structName string) (columns []system.SysTableColumns, err error) {
	err = global.HAB_DB.Where("struct_name = ?", structName).Preload("SysDataFilter").Find(&columns).Error
	return
}
