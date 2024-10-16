package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	ID   int64  `gorm:"column:id; type:int; primaryKey; autoIncrement; comment: 'PK is ID'"`
	Name string `gorm:"column:name"`
	Note string `gorm:"column:note; type:text"`
}

func (r *Role) TableName() string {
	return "roles"
}
