package utils

import "gin_vue_blog_server/model/response"

func VertifyResMsgType(msgType response.MsgType, msgCode response.MsgStrKey, defaultErr string) string {
	var ok bool
	var msg string
	switch msgType {
	case response.ImageMsg:
		msg, ok = response.ImageMsgMap[msgCode]
	case response.FileMsg:
		msg, ok = response.FileMsgMap[msgCode]
	case response.SucMsg:
		msg, ok = response.SuccessMsgMap[msgCode]
	default:
		ok = false
	}

	if !ok {
		return defaultErr
	}

	return msg
}
