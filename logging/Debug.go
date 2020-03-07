//+build debug

package logging

import (
	"log"
	"os"
	"strings"
)

func Info(format string, v ...interface{}) {
	f, err := os.OpenFile("native-shell.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf(format, v...)
}

func ProcessStr(str string, maxLen int) string {
	if len(str) > maxLen {
		str = str[:maxLen] + "..."
	}
	return strings.ReplaceAll(str, "\n", "//")
}