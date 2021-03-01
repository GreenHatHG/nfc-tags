package main

import (
	_ "nfc-tags/boot"
	_ "nfc-tags/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
