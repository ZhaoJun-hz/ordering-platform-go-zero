// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRoleMenu = "sys_role_menu"

// SysRoleMenu mapped from table <sys_role_menu>
type SysRoleMenu struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID int64 `gorm:"column:role_id;not null" json:"role_id"`
	MenuID int64 `gorm:"column:menu_id;not null" json:"menu_id"`
}

// TableName SysRoleMenu's table name
func (*SysRoleMenu) TableName() string {
	return TableNameSysRoleMenu
}