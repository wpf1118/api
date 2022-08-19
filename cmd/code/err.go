package code

import "github.com/wpf1118/toolbox/tools/errno"

var (
	ParseParamError    = errno.NewError(10000, "参数解析错误")
	ParseParamRequired = errno.NewError(10000, "参数%s不能为空")
	ParseParamInvalid  = errno.NewError(10000, "参数%s不合法")
	UserNotExists      = errno.NewError(20000, "该用户不存在或已被删除")
)
