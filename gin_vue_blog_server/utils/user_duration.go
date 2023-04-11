package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(dr string) (time.Duration, error) {
	dr = strings.TrimSpace(dr)
	duration, err := time.ParseDuration(dr)
	if err == nil {
		return duration, err
	}
	if strings.Contains(dr, "d") {
		dIndex := strings.Index(dr, "d")
		hour, _ := strconv.Atoi(dr[:dIndex])
		duration = time.Hour * 24 * time.Duration(hour)
		nDuration, err := time.ParseDuration(dr[dIndex+1:])
		if err != nil {
			return duration, nil
		}
		return duration + nDuration, nil
	}
	dv, err := strconv.ParseInt(dr, 10, 64)
	return time.Duration(dv), err
}
