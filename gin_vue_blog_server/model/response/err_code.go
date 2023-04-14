package response

type ErrCode int

const (
	SettingError      ErrCode = 3001
	FailNotExistError ErrCode = 3002
)

var (
	ErrMap = map[ErrCode]string{
		FailNotExistError: "文件不存在！",
		SettingError:      "系统配置出错",
	}
)
