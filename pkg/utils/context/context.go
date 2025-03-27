package context

import (
	"context"
	"encoding/json"
	"fmt"
	"ordering-platform/pkg/xerr"
)

// GetRoleIDFromContext 从上下文中获取并转换 roleId (string → int64)
func GetRoleIDFromContext(ctx context.Context) (int64, error) {
	// 1. 检查上下文中是否存在 roleId
	val := ctx.Value("roleId")
	if val == nil {
		return 0, xerr.NewCodeInvalidArgumentError("failed to get roleId from context")
	}

	roleId, ok := val.(json.Number)
	if !ok {
		return 0, xerr.NewCodeInvalidArgumentError(fmt.Sprintf("roleId is not a number (actual type: %T)", val))
	}

	number, err := roleId.Int64()
	if err != nil {
		return 0, xerr.NewCodeInvalidArgumentError(fmt.Sprintf("roleId is not a number"))
	}

	// 4. （可选）验证业务逻辑，如非负数
	if number < 0 {
		return 0, xerr.NewCodeInvalidArgumentError("roleId cannot be negative")
	}

	return number, nil
}

// GetUserIDFromContext 从上下文中获取并转换 userId (string → int64)
func GetUserIDFromContext(ctx context.Context) (int64, error) {
	// 1. 检查上下文中是否存在 userId
	val := ctx.Value("userId")
	if val == nil {
		return 0, xerr.NewCodeInvalidArgumentError("failed to get userId from context")
	}

	userId, ok := val.(json.Number)
	if !ok {
		return 0, xerr.NewCodeInvalidArgumentError(fmt.Sprintf("userId is not a string (actual type: %T)", val))
	}

	number, err := userId.Int64()
	if err != nil {
		return 0, xerr.NewCodeInvalidArgumentError(fmt.Sprintf("userId is not a number"))
	}

	// 4. （可选）验证业务逻辑，如非负数
	if number < 0 {
		return 0, xerr.NewCodeInvalidArgumentError("userId cannot be negative")
	}

	return number, nil
}
