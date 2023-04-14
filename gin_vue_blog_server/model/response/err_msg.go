package response

// 错误信息对应表
type MsgType int
type MsgStrKey int

const (
	SucMsg   = 0
	ErrMsg   = 9
	ImageMsg = 4
	FileMsg  = 3

	SuccessMsg         = 0000
	ImageUploadSuccess = 0001
	ErrorMsg           = 9999
	FileUploadMd5Err   = 3001
	ImageUploadTypeErr = 4001
	ImageUploadSizeErr = 4002
)

var (
	ImageMsgMap = map[MsgStrKey]string{
		ImageUploadTypeErr: "上传失败，图片文件类型不匹配！",
		ImageUploadSizeErr: "上传失败，图片大小超过限制大小",
	}

	FileMsgMap = map[MsgStrKey]string{
		FileUploadMd5Err: "上传失败，数据库中已有相同的文件！",
	}

	SuccessMsgMap = map[MsgStrKey]string{
		SuccessMsg:         "操作成功！",
		ImageUploadSuccess: "图片上传成功！",
	}
)
