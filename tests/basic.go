package main

import (
	"ansicode"
	"os"
	"runtime"
)

func Fatal(msg string) {
	txt := ansicode.Span("error").Red().Span(" " + msg + "\n")
	os.Stdout.WriteString(txt.String())
}

func main() {

	func runtime_pollServerInit()
	runtime_pollServerInit()

	Fatal("must be root")

}