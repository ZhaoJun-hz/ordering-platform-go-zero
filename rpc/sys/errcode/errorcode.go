package errcode

import "ordering-platform/pkg/xerr"

// rpc 全部是20xxxxx错误码
// 01 是sys rpc 错误码
var (
	UserPwdError                = xerr.New(2001001, "用户名或密码不正确")
	UserStatusError             = xerr.New(2001002, "用户已被禁用")
	UserNotExistError           = xerr.New(2010003, "用户不存在")
	RoleNotExistError           = xerr.New(2010004, "角色不存在")
	ApiSelectError              = xerr.New(2010005, "选择的API列表错误")
	MenuTypeError               = xerr.New(2010006, "菜单类型错误")
	MenuNotExistError           = xerr.New(2010007, "菜单不存在")
	MenuUpdateMenuTypeError     = xerr.New(2010008, "更新菜单时，不允许修改菜单类型")
	MenuHaveAllocationRoleError = xerr.New(2010009, "菜单已经分配过角色，不允许修改")
	MenuHaveSubMenuError        = xerr.New(2010010, "菜单下还有子菜单，不允许删除")
)
