// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	UserID      int64          `gorm:"column:user_id;primaryKey;autoIncrement:true;comment:编码" json:"user_id"`              // 编码
	Status      int32          `gorm:"column:status;not null;default:1;comment:状态 1 正常 2 禁用" json:"status"`                 // 状态 1 正常 2 禁用
	Username    string         `gorm:"column:username;not null;comment:登录名" json:"username"`                                // 登录名
	Password    string         `gorm:"column:password;not null;comment:密码" json:"password"`                                 // 密码
	Nickname    string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`                                 // 昵称
	Description string         `gorm:"column:description;not null;comment:用户的描述信息" json:"description"`                      // 用户的描述信息
	Mobile      string         `gorm:"column:mobile;not null;comment:手机号" json:"mobile"`                                    // 手机号
	Email       string         `gorm:"column:email;not null;comment:邮箱号" json:"email"`                                      // 邮箱号
	Avatar      string         `gorm:"column:avatar;not null;comment:头像路径" json:"avatar"`                                   // 头像路径
	DeptID      int64          `gorm:"column:dept_id;not null;default:1;comment:部门ID" json:"dept_id"`                       // 部门ID
	RoleID      int64          `gorm:"column:role_id;not null;default:1;comment:角色ID" json:"role_id"`                       // 角色ID
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   *time.Time     `gorm:"column:updated_at;comment:最后更新时间" json:"updated_at"`                                  // 最后更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	CreateBy    int64          `gorm:"column:create_by;not null;default:1;comment:创建者" json:"create_by"`                    // 创建者
	UpdateBy    int64          `gorm:"column:update_by;not null;default:1;comment:更新者" json:"update_by"`                    // 更新者
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
