package errcode

import "github.com/wpf1118/toolbox/tools/errno"

var (
	ParseParamError    = errno.NewError(10000, "参数解析错误")
	ParseParamRequired = errno.NewError(10001, "参数%s不能为空")
	ParseParamInvalid  = errno.NewError(10002, "参数%s不合法")
	GetDataError       = errno.NewError(10003, "查询数据出错")

	DataNotExists       = errno.NewError(20000, "该数据不存在或已被删除")
	UserNotExists       = errno.NewError(20001, "该用户不存在或已被删除")
	CommonSetError      = errno.NewError(20002, "保存失败")
	FromFileError       = errno.NewError(20003, "获取上传文件失败")
	CreateFileError     = errno.NewError(20004, "创建文件失败")
	CopyFileError       = errno.NewError(20005, "拷贝文件失败")
	CreateDirError      = errno.NewError(20006, "创建目录失败")
	FileExists          = errno.NewError(20007, "该文件已存在")
	FileNotExists       = errno.NewError(20008, "该文件不存在或已被删除")
	NotSupportedFileExt = errno.NewError(20009, "暂不支持该文件类型")
	FileTooLarge        = errno.NewError(20010, "上传文件过大")

	// category

)
