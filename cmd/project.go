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
	module := &Module{
		ModName: "app",
	}
	if _, err := os.Stat(module.ModName); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(module.ModName, 0754); err != nil {
			return err
		}
	}

	err := generateModule(module)
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
