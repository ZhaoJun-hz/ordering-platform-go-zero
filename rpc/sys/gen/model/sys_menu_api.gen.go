// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysMenuAPI = "sys_menu_api"

// SysMenuAPI mapped from table <sys_menu_api>
type SysMenuAPI struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MenuID int64 `gorm:"column:menu_id;not null" json:"menu_id"`
	APIID  int64 `gorm:"column:api_id;not null;comment:主键编码" json:"api_id"` // 主键编码
}

// TableName SysMenuAPI's table name
func (*SysMenuAPI) TableName() string {
	return TableNameSysMenuAPI
}
