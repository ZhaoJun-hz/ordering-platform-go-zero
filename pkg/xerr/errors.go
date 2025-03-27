/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package xerr

import "github.com/zeromicro/x/errors"

func New(code int, msg string) error {
	return errors.New(code, msg)
}

func NewMsg(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

func NewDBErr() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}

// NewCodeInvalidArgumentError returns Code Error with custom invalid argument error code
func NewCodeInvalidArgumentError(msg string) error {
	return errors.New(REQUEST_PARAM_ERROR, msg)
}
