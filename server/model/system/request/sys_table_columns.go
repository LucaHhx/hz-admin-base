package request

import (
	"hz-admin-base/model/common/request"
	"time"
)

type SysTableColumnsSearch struct {
	MenuId         *uint      `json:"menuId" form:"menuId"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StructName     *string    `json:"structName" form:"structName" `
	JsonName       *string    `json:"jsonName" form:"jsonName" `
	ColumnName     *string    `json:"columnName" form:"columnName" `
	StartWith      *int32     `json:"startWith" form:"startWith"`
	EndWith        *int32     `json:"endWith" form:"endWith"`
	Type           *string    `json:"type" form:"type" `
	Sortable       *string    `json:"sortable" form:"sortable" `
	Filter         *string    `json:"filter" form:"filter" `
	DefaultFilter  *string    `json:"defaultFilter" form:"defaultFilter" `
	FilterList     *string    `json:"filterList" form:"filterList" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
