package utils

import (
	"fmt"
	"hz-admin-base/global"
	"hz-admin-base/model/common/request"
	"hz-admin-base/model/system"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
)

type Table interface {
	TableName() string
}

var ColumnsCache = cmap.NewWithCustomShardingFunction[uint, map[string][]*system.SysAuthorityCol](func(key uint) uint32 {
	return uint32(key)
})

func GetTbName(db *gorm.DB) string {
	tableName := db.Statement.Table
	if tableName == "" && db.Statement.Model != nil {
		if tb, ok := db.Statement.Model.(Table); ok {
			tableName = tb.TableName()
		}
	}
	return tableName
}

func FindTableColumns(tableName string, aid uint) []*system.SysAuthorityCol {
	if tbcols, ok := ColumnsCache.Get(aid); ok {
		if cols, ok := tbcols[tableName]; ok {
			return cols
		}
	}
	var authority []*system.SysAuthorityCol
	err := global.GVA_DB.Model(system.SysAuthorityCol{}).
		Where("authority_id = ?", aid).
		Preload("SysTableColumns").Find(&authority).Error
	if err != nil {
		return nil
	}
	colGroup := lo.GroupBy(authority, func(item *system.SysAuthorityCol) string {
		return item.SysTableColumns.TbName
	})
	ColumnsCache.Set(aid, colGroup)
	return colGroup[tableName]
}

func GetColumns(db *gorm.DB, c *gin.Context) (columns []string) {
	authorityId := GetUserAuthorityId(c)
	tableName := GetTbName(db)
	cols := FindTableColumns(tableName, authorityId)
	if len(cols) == 0 {
		return []string{"*"}
	}
	keys := lo.Map(cols, func(item *system.SysAuthorityCol, index int) string {
		if item.SysTableColumns.IsAdditional || item.SysTableColumns.IsHaving || !item.SysTableColumns.IsShow {
			return ""
		}
		return fmt.Sprintf("`%s`.`%s`", tableName, item.SysTableColumns.ColumnName)
	})
	return lo.Filter(keys, func(item string, index int) bool {
		return strings.TrimSpace(item) != ""
	})
}

func GetUpdateColumns(db *gorm.DB, c *gin.Context) (columns []string) {
	authorityId := GetUserAuthorityId(c)
	tableName := GetTbName(db)
	cols := FindTableColumns(tableName, authorityId)
	if len(cols) == 0 {
		return []string{"*"}
	}
	gvaColumns := []string{"id", "created_at", "updated_at", "deleted_at", "created_by", "updated_by", "deleted_by"}
	keys := lo.Map(cols, func(item *system.SysAuthorityCol, index int) string {
		if item.SysTableColumns.IsAdditional || item.SysTableColumns.IsHaving {
			return ""
		}
		if (item.SysTableColumns.FormDisabled || item.SysTableColumns.FormHidden) && !utils.Contains(gvaColumns, item.SysTableColumns.ColumnName) {
			return ""
		}
		return fmt.Sprintf("`%s`.`%s`", tableName, item.SysTableColumns.ColumnName)
	})
	return lo.Filter(keys, func(item string, index int) bool {
		return strings.TrimSpace(item) != ""
	})
}

func GetColumnsMap(tableName string, aid uint) map[string]*system.SysAuthorityCol {
	cols := FindTableColumns(tableName, aid)
	data := make(map[string]*system.SysAuthorityCol)
	for _, col := range cols {
		data[col.SysTableColumns.ColumnName] = col
	}
	return data
}

func TableQuery(db *gorm.DB, info request.QueryInfo, c *gin.Context) (count int64, err error) {
	tableName := GetTbName(db)
	var colMap map[string]*system.SysAuthorityCol
	// 搜索条件处理
	if info.SearchInfo != nil {
		authorityId := GetUserAuthorityId(c)
		colMap = GetColumnsMap(tableName, authorityId)
		for key, value := range info.SearchInfo {
			if value == nil || value.Value == "" || colMap[key] == nil || colMap[key].SysTableColumns.IsAdditional {
				continue
			}
			tbaleKey := fmt.Sprintf("`%s`.`%s`", tableName, key)
			switch value.Option {
			case "in":
				// 将逗号分隔的字符串转为数组
				values := strings.Split(value.Value, ",")
				db = db.Where(tbaleKey+" IN ?", values)
			case "not in":
				// 将逗号分隔的字符串转为数组
				values := strings.Split(value.Value, ",")
				db = db.Where(tbaleKey+" NOT IN ?", values)
			case "between":
				// 将逗号分隔的字符串转为两个值
				values := strings.Split(value.Value, ",")
				if len(values) == 2 {
					db = db.Where(fmt.Sprintf("%s >= ? and %s < ?", tbaleKey, tbaleKey), values[0], values[1])
				}
			case "not between":
				// 将逗号分隔的字符串转为两个值
				values := strings.Split(value.Value, ",")
				if len(values) == 2 {
					db = db.Where(fmt.Sprintf("%s < ? or %s >= ?", tbaleKey, tbaleKey), values[0], values[1])
				}
			case "like":
				// 处理模糊查询
				db = db.Where(tbaleKey+" LIKE ?", "%"+value.Value+"%")
			default:
				db = db.Where(tbaleKey+" "+value.Option+" ?", value.Value)
			}
		}
	}
	// 排序处理
	if info.SortInfo != nil {
		// 创建排序切片以保持顺序
		type sortItem struct {
			key   string
			order string
			index int
		}
		sortItems := make([]sortItem, 0)

		// 收集排序信息
		for key, value := range info.SortInfo {
			if value.Type != "" {
				sortItems = append(sortItems, sortItem{
					key:   fmt.Sprintf("`%s`.`%s`", tableName, key),
					order: value.Type, // 'asc' 或 'desc'
					index: value.Index,
				})
			}
		}

		// 按照index排序
		sort.Slice(sortItems, func(i, j int) bool {
			return sortItems[i].index < sortItems[j].index
		})

		// 应用排序
		for _, item := range sortItems {
			db = db.Order(item.key + " " + item.order)
		}
	}
	// 获取总条数
	err = db.Count(&count).Error

	// 分页处理
	pageInfo := info.PageInfo
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	TableHaving(db, info, colMap)
	db = db.Limit(limit).Offset(offset)
	return
}

func TableHaving(db *gorm.DB, info request.QueryInfo, colMap map[string]*system.SysAuthorityCol) {
	if info.SearchInfo != nil {
		for key, value := range info.SearchInfo {
			if value == nil || value.Value == "" || colMap[key] == nil || !colMap[key].SysTableColumns.IsAdditional {
				continue
			}
			tbaleKey := fmt.Sprintf("`%s`", key)
			switch value.Option {
			case "in":
				// 将逗号分隔的字符串转为数组
				values := strings.Split(value.Value, ",")
				db = db.Having(tbaleKey+" IN ?", values)
			case "not in":
				// 将逗号分隔的字符串转为数组
				values := strings.Split(value.Value, ",")
				db = db.Having(tbaleKey+" NOT IN ?", values)
			case "between":
				// 将逗号分隔的字符串转为两个值
				values := strings.Split(value.Value, ",")
				if len(values) == 2 {
					db = db.Having(fmt.Sprintf("%s >= ? and %s < ?", tbaleKey, tbaleKey), values[0], values[1])
				}
			case "not between":
				// 将逗号分隔的字符串转为两个值
				values := strings.Split(value.Value, ",")
				if len(values) == 2 {
					db = db.Having(fmt.Sprintf("%s < ? or %s >= ?", tbaleKey, tbaleKey), values[0], values[1])
				}
			default:
				db = db.Having(tbaleKey+" "+value.Option+" ?", value.Value)
			}
		}
	}
}
