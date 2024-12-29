package errorMsg

import "github.com/fatih/color"

func NotFoundCmd() {
    color.Red("command not found.")
}

func NotImplemented() {
    color.Red("command not found.")
}

func UnknownOption(option string) {
    color.Red("unknown option:", option)
}