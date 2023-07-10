package errex

var (
	ErrOpenIdEmpty     = Create(-1, "账号获取失败")
	ErrOpenIdGetErr    = Create(-1, "账号信息获取异常")
	ErrUnknow          = Create(-1, "未知异常")
	ErrServerListEmpty = Create(-1, "服务器列表为空")
	ErrLoginKeyError   = Create(10001, "Login verification failed")
)
