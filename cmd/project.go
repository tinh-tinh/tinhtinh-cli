package cmd

import (
	"os"
	"text/template"

	"github.com/tinh-tinh/tinhtinh-cli/tpl"
)

type Project struct {
	// v2
	PkgName string
}

type Command struct {
	CmdName   string
	CmdParent string
	*Project
}

func (p *Project) Create() error {
	err := generateModule(&Module{
		ModName: "app",
	})
	if err != nil {
		return err
	}

	appFile, err := os.Create("app/app_module.go")
	if err != nil {
		return err
	}
	defer appFile.Close()

	appTemplate := template.Must(template.New("app").Parse(string(tpl.AppTemplate())))
	err = appTemplate.Execute(appFile, nil)
	if err != nil {
		return err
	}

	mainFile, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	err = mainTemplate.Execute(mainFile, p)
	if err != nil {
		return err
	}

	return nil
}
