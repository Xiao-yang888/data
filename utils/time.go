package utils

import "time"

/*
 *2006年01月02日 15:04:05
 *2006/01/02 15:04:05
 *2006-01-02 15:04:05
 *2006.01.02日 15:04:05
 */
const TIME_Format_ONE = "2006年01月02日 15:04:05"
const TIME_Format_TWO = "2006/01/02 15:04:05"
const TIME_Format_THREE = "2006-01-02 15:04:05"
const TIME_Format_FOUR = "2006.01.02日 15:04:05"

/*
 *时间的格式化操作
 */
func TimeFormat(t int64, format string) string {
	return  time.Unix(t, 0).Format("format")
}

