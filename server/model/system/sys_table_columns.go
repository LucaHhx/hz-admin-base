// 自动生成模板SysTableColumns
package system

import (
	"hab/model/gtype"
)

// 表字段配置 结构体  SysTableColumns
type SysTableColumns struct {
	ID             uint              `gorm:"primarykey" json:"ID"`                                                                                    // 主键ID
	TbName         string            `json:"tbName" form:"tbName" gorm:"uniqueIndex:idx_table_col;column:tb_name;comment:表名称;size:255;"`              // 表名称
	MenuId         uint              `json:"menuId" form:"menuId" gorm:"uniqueIndex:idx_table_col;column:menu_id;comment:菜单ID;type:int;"`             // 菜单ID
	StructName     string            `json:"structName" form:"structName" gorm:"uniqueIndex:idx_table_col;column:struct_name;comment:结构名称;size:255;"` //结构名称
	JsonName       string            `json:"jsonName" form:"jsonName" gorm:"uniqueIndex:idx_table_col;column:json_name;comment:json名称;size:255;"`     //json名称
	ColumnName     string            `json:"columnName" form:"columnName" gorm:"column:column_name;comment:数据库名称;size:255;"`                          //数据库名称
	With           int32             `json:"with" form:"with" gorm:"column:with;comment:宽度;type:int;"`                                                //宽度
	Type           string            `json:"type" form:"type" gorm:"column:type;comment:类型;size:255"`                                                 //类型
	Sortable       bool              `json:"sortable" form:"sortable" gorm:"column:sortable;comment:排序;type:tinyint;"`                                //排序
	Filter         bool              `json:"filter" form:"filter" gorm:"column:filter;comment:筛选;type:tinyint;"`                                      //筛选
	DefaultFilter  string            `json:"defaultFilter" form:"defaultFilter" gorm:"column:default_filter;comment:默认筛选;size:255;"`                  //默认筛选
	FilterList     gtype.Arr[string] `json:"filterList" form:"filterList" gorm:"column:filter_list;comment:筛选列表;type:json;"`                          //筛选列表
	Sort           int               `json:"sort" form:"sort" gorm:"column:sort;comment:排序"`
	Note           string            `json:"note" form:"note" gorm:"column:note;comment:备注;type:text;"`                                     //备注
	IsShow         bool              `json:"isShow" form:"isShow" gorm:"column:is_show;comment:是否显示;type:tinyint;"`                         //是否显示
	Enum           gtype.Arr[string] `json:"enum" form:"enum" gorm:"column:enum;comment:枚举;type:json;"`                                     //枚举
	Fixed          string            `json:"fixed" form:"fixed" gorm:"column:fixed;comment:固定;size:255;"`                                   //固定
	Menu           SysBaseMenu       `json:"menu" gorm:"foreignKey:MenuId;references:ID"`                                                   //菜单
	Format         string            `json:"format" form:"format" gorm:"column:format;comment:格式;size:255;"`                                //格式
	IsAddSearch    bool              `json:"isAddSearch" form:"isAddSearch" gorm:"column:is_add_search;comment:是否添加搜索;type:tinyint;"`       //是否添加搜索
	SearchWidth    int32             `json:"searchWidth" form:"searchWidth" gorm:"column:search_width;comment:搜索宽度;type:int;"`              //搜索宽度
	IsAdditional   bool              `json:"isAdditional" form:"isAdditional" gorm:"column:is_additional;comment:是否附加;type:tinyint;"`       //是否附加
	IsHaving       bool              `json:"isHaving" form:"isHaving" gorm:"column:is_having;comment:是否having;type:tinyint;"`               //是否having
	FilterId       uint              `json:"filterId" form:"filterId" gorm:"column:filter_id;comment:筛选ID;type:int;"`                       //筛选ID
	EditorFilterId uint              `json:"editorFilterId" form:"editorFilterId" gorm:"column:editor_filter_id;comment:编辑器筛选ID;type:int;"` //编辑器筛选ID
	CreatedBy      uint              `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint              `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint              `gorm:"column:deleted_by;comment:删除者"`
	SysDataFilter  SysDataFilter     `json:"sysDataFilter" gorm:"foreignKey:FilterId;references:ID"` //筛选ID
	FormInfo
	IsSum bool `json:"isSum" form:"isSum" gorm:"column:is_sum;comment:是否求和;type:tinyint;"`
}

type FormInfo struct {
	FormWith     int  `json:"formWith" form:"formWith" gorm:"column:form_with;comment:表单宽度;type:int;"`                 //表单宽度
	FormDisabled bool `json:"formDisabled" form:"formDisabled" gorm:"column:form_disabled;comment:表单禁用;type:tinyint;"` //表单禁用
	FormHidden   bool `json:"formHidden" form:"formHidden" gorm:"column:form_hidden;comment:表单隐藏;type:tinyint;"`       //表单隐藏
	FormOrder    int  `json:"formOrder" form:"formOrder" gorm:"column:form_order;comment:表单排序;type:int;"`              //表单排序
	FormMust     bool `json:"formMust" form:"formMust" gorm:"column:form_must;comment:表单必填;type:tinyint;"`             //表单必填
}

// TableName 表字段配置 SysTableColumns自定义表名 sys_table_columns
func (SysTableColumns) TableName() string {
	return "sys_table_columns"
}
