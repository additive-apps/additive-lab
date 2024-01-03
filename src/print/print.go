package print

import (
	"fmt"
	"os"
)

var ansiColors = ANSIColors()
var unicodeIcons = UnicodeIcons()

func Done(format string, message ...any) {
	print(format, unicodeIcons.done, "", message...)
}

func Looking(format string, message ...any) {
	print(format, unicodeIcons.looking, "", message...)
}

func Fetching(format string, message ...any) {
	print(format, unicodeIcons.fetching, "", message...)
}

func Linking(format string, message ...any) {
	print(format, unicodeIcons.fetching, "", message...)
}

func Building(format string, message ...any) {
	print(format, unicodeIcons.building, "", message...)
}

func Normal(format string, message ...any) {
	fmt.Printf(format, message...)
}

func Info(format string, message ...any) {
	print(format, "info", ansiColors.green, message...)
}

func Create(format string, message ...any) {
	print(format, "create", ansiColors.cyan, message...)
}

func Success(format string, message ...any) {
	print(format, "success", ansiColors.green, message...)
}

func Warning(format string, message ...any) {
	print(format, "warning", ansiColors.yellow, message...)
}

func Fatal(format any, message ...any) {
	print(format, "fatal", ansiColors.red, message...)
	os.Exit(1)
}

func print(format any, printType string, printTypeColor string, message ...any) {
	fmt.Printf(
		formater(printTypeColor, printType, ansiColors.systemDefault, format),
		message...,
	)
}

func formater(printTypeColor string, printType string, textColor string, text any) string {
	return fmt.Sprintf(
		"%s%s%s %s\n",
		printTypeColor,
		printType,
		textColor,
		text.(string),
	)
}
