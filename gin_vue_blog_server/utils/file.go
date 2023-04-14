package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gin_vue_blog_server/model/response"
	"os"
)

func IsPathExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if fileInfo.IsDir() {
		return true, nil
	}
	return false, errors.New("存在同名文件")
}

func CheckStrInList(str string, whiteList []string) bool {
	for _, whiteStr := range whiteList {
		if whiteStr == str {
			return true
		}
	}
	return false
}

func InitFileResponse(filename string, code int, msgCode response.MsgStrKey,
	msgType response.MsgType, err error) response.FileResponse {

	var msg string

	if msgType == response.ErrMsg {
		msg = err.Error()
	} else {
		msg = VertifyResMsgType(msgType, msgCode, "文件上传失败！")
	}

	return response.FileResponse{
		FileName: filename,
		Code:     code,
		Msg:      msg,
	}
}

func InitMd5Str(fileStr string) string {
	data := []byte(fileStr)
	hash := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", hash)
	return md5Str
}
