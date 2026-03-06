package global

import (
	"gorm.io/gorm"
)

type HAB_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt MySQLTime      // 创建时间
	UpdatedAt MySQLTime      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type HAB_MODEL_NOD struct {
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt MySQLTime // 创建时间
	UpdatedAt MySQLTime // 更新时间
}

type HAB_MODEL_NOID struct {
	CreatedAt MySQLTime      // 创建时间
	UpdatedAt MySQLTime      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
