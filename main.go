package main

import (
	"fmt"
	"tsumikey/pkg/misskey"
	"tsumikey/pkg/tui"
)

func main() {
	fmt.Println("test")
	misskey.API_Test()
	tui.TUI_Test()

}
