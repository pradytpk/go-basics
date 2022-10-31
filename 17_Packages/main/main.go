package main

import (
	"example.com/code_samples/ztm/go-basics/17_Packages/display"
	"example.com/code_samples/ztm/go-basics/17_Packages/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello pradeep")
	msg.Exiting("test")
}
