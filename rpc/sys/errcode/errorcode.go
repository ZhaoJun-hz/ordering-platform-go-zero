package errcode

import "ordering-platform/pkg/xerr"

// rpc 全部是20xxxxx错误码
// 01 是sys rpc 错误码
var (
	UserPwdError                = xerr.New(20010001, "用户名或密码不正确")
	UserStatusError             = xerr.New(20010002, "用户已被禁用")
	UserNotExistError           = xerr.New(20010003, "用户不存在")
	RoleNotExistError           = xerr.New(20010004, "角色不存在")
	ApiSelectError              = xerr.New(20010005, "选择的API列表错误")
	MenuTypeError               = xerr.New(20010006, "菜单类型错误")
	MenuNotExistError           = xerr.New(20010007, "菜单不存在")
	MenuUpdateMenuTypeError     = xerr.New(20010008, "更新菜单时，不允许修改菜单类型")
	MenuHaveAllocationRoleError = xerr.New(20010009, "菜单已经分配过角色，不允许修改")
	MenuHaveSubMenuError        = xerr.New(20010010, "菜单下还有子菜单，不允许删除")

	DeptNotExistError    = xerr.New(20010030, "部门不存在")
	DeptNotEmptyError    = xerr.New(20100031, "部门还有用户")
	DeptHaveSubDeptError = xerr.New(20100032, "部门下还有子部门")
)
