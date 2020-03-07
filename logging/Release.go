//+build !debug

package logging

func Info(format string, v ...interface{}) {}

func ProcessStr(str string, maxLen int) string { return str }