package logUtil

import (
	"fmt"
	"reflect"
)

const (
	// itoa在新的const块内会变为0, 可以通过+一个数值使const field内的变量保持递增
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

func Black(str string) string {
	return textColor(textBlack, str)
}

func Red(str string) string {
	return textColor(textRed, str)
}
func Yellow(str string) string {
	return textColor(textYellow, str)
}
func Green(str string) string {
	return textColor(textGreen, str)
}
func Cyan(str string) string {
	return textColor(textCyan, str)
}
func Blue(str string) string {
	return textColor(textBlue, str)
}
func Purple(str string) string {
	return textColor(textPurple, str)
}
func White(str string) string {
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}


func GetStructureDataInfo(itf interface{}, structName string) string {

	structType 	:= reflect.TypeOf(itf)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		panic("can not use not-structure arguments" +
				 " in function \"GetStructureDataInfo\"")
	}

	structValue	:= reflect.ValueOf(itf)

	info := fmt.Sprintf("\t%s {\n", Green(structName))

	for i := 0; i < structType.NumField(); i++ {
		typeInfo := structType.Field(i).Name
		valueInfo := structValue.Field(i)

		info += fmt.Sprintf("\t\t%-30s: %v\n", typeInfo, valueInfo)
	}
	info += "\t}"
	return info
}
