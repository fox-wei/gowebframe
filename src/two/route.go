package main

import (
	"github.com/fox-wei/gowebframe/framework"
)

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler) //*设置控制器
}
