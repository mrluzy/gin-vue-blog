package res

type (
	ErrorCode int
	ErrorMap  map[ErrorCode]string
)

const (
	SettingsError ErrorCode = 1001
)

var (
	ErrMap = ErrorMap{
		1001: "系统错误",
	}
)
