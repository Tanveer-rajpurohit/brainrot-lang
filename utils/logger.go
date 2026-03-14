package utils

import (
	"fmt"
)

const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue   = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan   = "\033[36m"
    ColorBold   = "\033[1m"
)

// Info prints a blue info message
func Info(msg string) {
    fmt.Printf("%s[INFO]%s %s\n", ColorBlue, ColorReset, msg)
}

// Success prints a green success message
func Success(msg string) {
    fmt.Printf("%s[OK]%s %s\n", ColorGreen, ColorReset, msg)
}

// Warn prints a yellow warning
func Warn(msg string) {
    fmt.Printf("%s[WARN]%s %s\n", ColorYellow, ColorReset, msg)
}	

// Fatal prints a red error and exits
func Fatal(err error) {
    fmt.Printf("%s%s%s\n", ColorRed, err.Error(), ColorReset)
}


func Banner() {
    fmt.Printf(`%s
╔══════════════════════════════════════╗
║     🔥 BrainRot Lang v1.0 🔥         ║
║   the BrainRot programming language  ║
╚══════════════════════════════════════╝
%s`, ColorBold, ColorReset)
}