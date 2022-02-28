package main

import "github.com/hideaki10/coredemo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
