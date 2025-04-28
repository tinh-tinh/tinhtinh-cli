package tpl

func MainTemplate() []byte {
	return []byte(`
package main

import (
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"{{ .PkgName}}/app"
)

func main() {
	server := core.CreateFactory(app.NewModule)

	server.Listen(3000)
}
`)
}

func AppTemplate() []byte {
	return []byte(`
package app

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewModule() core.Module {
	appModule := core.NewModule(core.NewModuleOptions{
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewService},
	})

	return appModule
}
`)
}

func ModuleTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewModule(module core.Module) core.Module {
	{{ .ModName }}Module := module.New(core.NewModuleOptions{
		Controllers: []core.Controllers{NewController},
		Providers:   []core.Providers{NewService},
	})

	return {{ .ModName }}Module
}
	`)
}

func ControllerTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/v2/core"

func NewController(module core.Module) core.Controller {
	ctrl := module.NewController("{{ .ModName }}")

	ctrl.Post("", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Get("", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Get("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Put("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{"data": "ok"})
	})

	ctrl.Delete("{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{"data": "ok"})
	})

	return ctrl
}
	`)
}

func ServiceTemplate() []byte {
	return []byte(`
package {{ .ModName }}

import "github.com/tinh-tinh/tinhtinh/v2/core"

const {{ .UpperModName }}_SERVICE core.Provide = "{{ .UpperModName}}_SERVICE"

type {{ .ModName }}Service struct {}

func (s *{{ .ModName }}Service) Create(input interface{}) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Find() interface{} {
	return nil
}

func (s *{{ .ModName }}Service) FindById(id string) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Update(id string,input interface{}) interface{} {
	return nil
}

func (s *{{ .ModName }}Service) Delete(id string) interface{} {
	return nil
}

func NewService(module core.Module) core.Provider {
	svc := module.NewProvider(core.ProviderOptions{
		Name: {{ .UpperModName }}_SERVICE,
		Value: &{{ .ModName }}Service{},
	})

	return svc
}
	`)
}
