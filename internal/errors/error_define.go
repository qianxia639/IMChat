package errors

var (
	ParamsErr            = NewInvalidArgumentError("参数错误")
	ServerErr            = NewInternalError("服务异常")
	UnauthorizaedErr     = NewUnauthorizaedError("认证失败")
	AccountLockedErr     = NewError(PermissionDenied, "账户锁定,请稍后再试")
	UserNotFoundErr      = NewError(NotFound, "用户不存在")
	AccountExistErr      = NewError(AlreadyExists, "用户已存在")
	IncorrectPasswordErr = NewUnauthorizaedError("密码错误")
	RestPwdNotSame       = NewInvalidArgumentError("两次输入的密码不一致")
	NewPwdIsSame         = NewInvalidArgumentError("新旧密码不能一致")
	NewPwdBelowInterval  = NewInvalidArgumentError("两次修改密码的间隔不得低于7天")
	EmailCodeErr         = NewUnauthorizaedError("邮箱验证码错误")
	SendEmailCodeErr     = NewInternalError("发送邮件失败")
	EmailExistErr        = NewError(AlreadyExists, "邮箱已存在")
	PermissionDeniedErr  = NewError(PermissionDenied, "权限不足")
	NicknameExistErr     = NewInvalidArgumentError("用户昵称已存在")
)
