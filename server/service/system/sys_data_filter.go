package system

import (
	"database/sql"
	"fmt"
	"hz-admin-base/enum"
	"hz-admin-base/global"
	"hz-admin-base/model/common/request"
	"hz-admin-base/model/system"
	"hz-admin-base/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	//systemReq "hz-admin-base/model/system/request"
	"gorm.io/gorm"
)

type SysDataFilterService struct{}

// CreateSysDataFilter 创建sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) CreateSysDataFilter(sysDataFilter *system.SysDataFilter, c *gin.Context) (err error) {
	err = global.GVA_DB.Create(sysDataFilter).Error
	return err
}

// ImportSysDataFilter 导入sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) ImportSysDataFilter(tep enum.ImportType, list []system.SysDataFilter, c *gin.Context) (err error) {
	switch tep {
	case enum.ImportType_Full:
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 截断表（清空表并重置自增ID）
			if err := tx.Exec("TRUNCATE TABLE sys_data_filter").Error; err != nil {
				return err
			}
			// 批量写入
			if err := tx.Create(&list).Error; err != nil {
				return err
			}
			return nil
		})
	case enum.ImportType_Append:
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			for i, _ := range list {
				nowTime := global.NowMySQLTime()
				list[i].CreatedAt = nowTime
				list[i].UpdatedAt = nowTime
				list[i].ID = 0
			}
			// 批量写入
			if err := tx.Create(&list).Error; err != nil {
				return err
			}
			return nil
		})
	default:
		return errors.New("Invalid import type")
	}
}

// DeleteSysDataFilter 删除sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) DeleteSysDataFilter(ID string, userID uint, c *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysDataFilter{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&system.SysDataFilter{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSysDataFilterByIds 批量删除sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) DeleteSysDataFilterByIds(IDs []string, deleted_by uint, c *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.SysDataFilter{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&system.SysDataFilter{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSysDataFilter 更新sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) UpdateSysDataFilter(sysDataFilter system.SysDataFilter, c *gin.Context) (err error) {
	err = global.GVA_DB.Model(&system.SysDataFilter{}).Where("id = ?", sysDataFilter.ID).Updates(&sysDataFilter).Error
	return err
}

// GetSysDataFilter 根据ID获取sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) GetSysDataFilter(ID string, c *gin.Context) (sysDataFilter system.SysDataFilter, err error) {
	// err = global.GVA_DB.Where("id = ?", ID).First(&sysDataFilter).Error
	db := global.GVA_DB.Model(&system.SysDataFilter{})
	err = db.Where("id = ?", ID).First(&sysDataFilter).Error
	return
}

// GetSysDataFilterInfoList 分页获取sysDataFilter记录
// Author [yourname](https://github.com/yourname)
func (sysDataFilterService *SysDataFilterService) GetSysDataFilterInfoList(info request.QueryInfo, c *gin.Context) (list []system.SysDataFilter, total int64, err error) {
	// Create db
	db := global.GVA_DB.Model(&system.SysDataFilter{})
	var sysDataFilters []system.SysDataFilter
	total, err = utils.TableQuery(db, info, c)
	if err != nil {
		return
	}
	err = db.Select(utils.GetColumns(db, c)).Find(&sysDataFilters).Error
	return sysDataFilters, total, err
}
func (sysDataFilterService *SysDataFilterService) GetSysDataFilterPublic() {
	// This method is for getting data source defined data
	// Please implement it yourself
}

func (sysDataFilterService *SysDataFilterService) ExecuteSql(sqlStr string) (interface{}, error) {
	//data := make(map[string]any)
	db, err := global.GVA_DB.DB()
	if err != nil {
		return nil, err
	}
	// 执行查询但只拿结构
	rows, err := db.Query(sqlStr + " LIMIT 0") // LIMIT 0 可以避免取出数据
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 获取列类型（包括列名）
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	return lo.Map(columnTypes, func(item *sql.ColumnType, index int) system.SysDataFilterColumn {
		return system.SysDataFilterColumn{
			ColumnName: item.Name(),
			Label:      "",
			IsShow:     false,
			Filter:     false,
		}
	}), err
}

func (sysDataFilterService *SysDataFilterService) FilterData(filters []string, id uint) (interface{}, error) {
	var sysDataFilter system.SysDataFilter
	err := global.GVA_DB.Model(&system.SysDataFilter{}).Where("id = ?", id).First(&sysDataFilter).Error
	if err != nil {
		return nil, err
	}
	filterCols := lo.Filter(sysDataFilter.Columns, func(item system.SysDataFilterColumn, index int) bool {
		return item.Filter
	})

	havingStr := ""
	for _, filter := range filters {
		timFilter := strings.TrimSpace(filter)
		if timFilter == "" {
			continue
		}
		havingList := make([]string, 0)
		for _, col := range filterCols {
			havingList = append(havingList, fmt.Sprintf("%s LIKE '%%%s%%'", col.ColumnName, timFilter))
		}
		if len(havingList) > 0 {
			if havingStr != "" {
				havingStr += " AND "
			}
			havingStr += fmt.Sprintf("(%s)", strings.Join(havingList, " OR "))
		}
	}
	var data []map[string]interface{}
	if len(havingStr) > 0 {
		err = global.GVA_DB.Raw(fmt.Sprintf("%s HAVING %s", sysDataFilter.Sql, havingStr)).Scan(&data).Error
	} else {
		err = global.GVA_DB.Raw(sysDataFilter.Sql).Scan(&data).Error
	}
	return data, err
}
