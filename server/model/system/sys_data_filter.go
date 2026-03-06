// 自动生成模板SysDataFilter
package system

import (
	"hab/global"
	"hab/model/gtype"
)

// sysDataFilter 结构体  SysDataFilter
type SysDataFilter struct {
	global.HAB_MODEL
	Name      string                         `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`                                //名称
	Sql       string                         `json:"sql" form:"sql" gorm:"column:sql;comment:;type:text;"`                                  //SQL
	Columns   gtype.Arr[SysDataFilterColumn] `json:"columns" form:"columns" gorm:"column:columns;comment:;type:json;" swaggertype:"object"` //列
	Key       string                         `json:"key" form:"key" gorm:"column:key;comment:;size:255;"`                                   //主键
	Label     string                         `json:"label" form:"label" gorm:"column:label;comment:;size:255;"`                             //标题
	Value     string                         `json:"value" form:"value" gorm:"column:value;comment:;size:255;"`                             //值
	OrderBy   string                         `json:"orderBy" form:"orderBy" gorm:"column:order_by;comment:;size:255;"`                      //排序
	Note      string                         `json:"note" form:"note" gorm:"column:note;comment:;type:text;"`                               //备注
	CreatedBy uint                           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint                           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint                           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName sysDataFilter SysDataFilter自定义表名 sys_data_filter
func (SysDataFilter) TableName() string {
	return "sys_data_filter"
}

type SysDataFilterColumn struct {
	ColumnName string `json:"columnName" form:"columnName"` //列名
	Label      string `json:"label" form:"label"`           //列标题
	IsShow     bool   `json:"isShow" form:"isShow"`         //是否显示
	Filter     bool   `json:"filter" form:"filter"`         //是否过滤
	Sort       int    `json:"sort" form:"sort"`             //排序
}
