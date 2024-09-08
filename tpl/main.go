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

func ModuleTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/core"

func NewModule(module *core.DynamicModule) *core.DynamicModule {
	{{ .ModName }}Module := module.New(core.NewModuleOptions{
		Controllers: []core.Controller{NewController},
		Providers:   []core.Provider{NewService},
	})

	return {{ .ModName }}
}
	`)
}

func ControllerTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/core"

func NewController(module *core.DynamicModule) *core.DynamicController {
	ctrl := module.NewController("{{ .ModName }}")

		ctrl.Post("/", func(ctx core.Ctx) {
		ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Get("/", func(ctx core.Ctx) {
		ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Get("/{id}", func(ctx core.Ctx) {
		ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Put("/{id}", func(ctx core.Ctx) {
		ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Delete("/{id}", func(ctx core.Ctx) {
		ctx.JSON(core.Map{"data": "ok"})
	})

	return ctrl
}
	`)
}

func ServiceTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/core"

type {{ .ModName }}Service struct {}

func (s *{{ .ModName }}Service) Create(input interface{}) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Get(input interface{}) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Update(input interface{}) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Delete(input interface{}) interface{} {
	return nil
}

func NewService(module *core.DynamicModule) *core.Service {
	svc := module.NewProvider(&{{ .ModName }}Service{})

	return svc
}
	`)
}
