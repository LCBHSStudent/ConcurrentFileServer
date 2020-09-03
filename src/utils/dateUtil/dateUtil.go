package dateUtil

import "time"

const (
	FMTyyyyMMdd				= "2006-01-02"
	FMTyyyyMMddHHmmss		= "2006-01-02 15:04:05"
	FMTHHmmss				= "15:04:05"
	FMTyyyyMMddHHmmssss 	= "2006-01-02 15:04:05...000"
)

var defaultFormatString = FMTyyyyMMddHHmmssss

func SetDefaultFormatString(format string) {
	defaultFormatString = format
}

func GetCurFormatTime() string {
	curTime := time.Now()
	return curTime.Format(defaultFormatString)
}