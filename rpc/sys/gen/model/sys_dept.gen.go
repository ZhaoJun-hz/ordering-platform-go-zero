// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDept = "sys_dept"

// SysDept mapped from table <sys_dept>
type SysDept struct {
	DeptID    int64          `gorm:"column:dept_id;primaryKey;autoIncrement:true;comment:主键编码" json:"dept_id"`            // 主键编码
	ParentID  int64          `gorm:"column:parent_id;not null;comment:父级部门id" json:"parent_id"`                           // 父级部门id
	DeptPath  string         `gorm:"column:dept_path;not null;comment:部门路径 / 分割" json:"dept_path"`                        // 部门路径 / 分割
	DeptName  string         `gorm:"column:dept_name;not null;comment:部门名字" json:"dept_name"`                             // 部门名字
	Sort      int32          `gorm:"column:sort;not null;default:1;comment:排序" json:"sort"`                               // 排序
	Leader    string         `gorm:"column:leader;not null;comment:负责人" json:"leader"`                                    // 负责人
	Phone     string         `gorm:"column:phone;not null;comment:负责人手机号" json:"phone"`                                   // 负责人手机号
	Email     string         `gorm:"column:email;not null;comment:负责人邮箱" json:"email"`                                    // 负责人邮箱
	Status    int32          `gorm:"column:status;not null;default:2;comment:状态 1 启用  2 停用" json:"status"`                // 状态 1 启用  2 停用
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;comment:最后更新时间" json:"updated_at"`                                  // 最后更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	CreateBy  int64          `gorm:"column:create_by;not null;default:1;comment:创建者" json:"create_by"`                    // 创建者
	UpdateBy  int64          `gorm:"column:update_by;not null;default:1;comment:更新者" json:"update_by"`                    // 更新者
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
