package response

type ErrCode int

const (
	SettingError       ErrCode = 3001
	FailNotExistError  ErrCode = 3002
	ParamError         ErrCode = 3003
	ArticleDetailError ErrCode = 3004
)

var (
	ErrMap = map[ErrCode]string{
		FailNotExistError:  "文件不存在！",
		SettingError:       "系统配置出错",
		ParamError:         "参数获取错误",
		ArticleDetailError: "文章详情获取错误",
	}
)
