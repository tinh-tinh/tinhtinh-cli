package cmd

import (
	"os"
	"strings"
	"text/template"

	"github.com/tinh-tinh/tinhtinh-cli/v2/tpl"
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
	module := &Module{
		ModName:      "app",
		UpperModName: strings.ToUpper("app"),
	}
	if _, err := os.Stat(module.ModName); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(module.ModName, 0754); err != nil {
			return err
		}
	}

	err := generateService(module)
	if err != nil {
		return err
	}

	err = generateController(module)
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
