package localtime

import "time"

func Now() *time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	result := time.Now().In(loc)
	return &result
}
