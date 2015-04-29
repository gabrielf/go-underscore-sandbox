package main

import (
	"github.com/gabrielf/godep-sandbox/_demo/demopack"
	"github.com/gabrielf/godep-sandbox/demo/_demopack"
)

func main() {
	demopack.Foo()
	_demopack.Bar()
}
