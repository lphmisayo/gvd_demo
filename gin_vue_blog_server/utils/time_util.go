package utils

import "time"

func GetNowTimeHasFormat() time.Time {
	//nowTime := time.Now().Format("2006-01-02 15:04:05")
	nowTime := time.Now()
	return nowTime
}
