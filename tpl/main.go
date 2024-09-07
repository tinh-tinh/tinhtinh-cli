package tpl

func MainTemplate() []byte {
	return []byte(`
package main

import (
	"github.com/tinh-tinh/tinhtinh/core"
	"{{ .PkgName}}/app"
)

func main() {
	server := core.CreateFactory(app.NewModule, "api")

	server.Listen(3000)
}
`)
}

func AppTemplate() []byte {
	return []byte(`
package app

import (
	"github.com/tinh-tinh/tinhtinh/core"
)

func NewModule() *core.DynamicModule {
	appModule := core.NewModule(core.NewModuleOptions{
		Global: true,
	})

	return appModule
}
`)
}
