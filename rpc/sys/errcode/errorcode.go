package errcode

import "ordering-platform/pkg/xerr"

// rpc 全部是20xxxxx错误码
// 01 是sys rpc 错误码
var (
	UserPwdError      = xerr.New(2001001, "用户名或密码不正确")
	UserStatusError   = xerr.New(2001002, "用户已被禁用")
	UserNotExistError = xerr.New(2010003, "用户不存在")
	RoleNotExistError = xerr.New(2010004, "角色不存在")
)