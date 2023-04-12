package utils

import "github.com/mojocn/base64Captcha"

var Store = base64Captcha.DefaultMemStore

func InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
