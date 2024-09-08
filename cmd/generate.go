/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/tinh-tinh/tinhtinh-cli/tpl"
)

type Module struct {
	PkgName      string
	ModName      string
	AbsolutePath string
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use: "generate [type] [name]",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		var directive cobra.ShellCompDirective
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "ERROR: Choose type you want generate for your module")
			directive = cobra.ShellCompDirectiveDefault
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "ERROR: This command does not take any more arguments (but may accept flags)")
			directive = cobra.ShellCompDirectiveNoFileComp
		} else if len(args) == 2 {
			cmdType := args[0]
			standard := []string{"controller", "service", "module"}
			idx := slices.IndexFunc(standard, func(s string) bool { return cmdType == s })
			if idx == -1 {
				comps = cobra.AppendActiveHelp(comps, "ERROR: invalid command type")
				directive = cobra.ShellCompDirectiveNoFileComp
			}
		} else {
			comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
			directive = cobra.ShellCompDirectiveNoFileComp
		}
		return comps, directive
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("add needs a name for the command"))
		}

		wd, err := os.Getwd()
		cobra.CheckErr(err)

		fmt.Print(args)

		moduleName := validateCmdName(args[1])
		module := &Module{
			ModName:      moduleName,
			AbsolutePath: wd,
		}

		err = generateProject(module, args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Generated successfully 🚀")
	},
}

func generateProject(module *Module, typeCmd string) error {
	if _, err := os.Stat(module.ModName); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(module.ModName, 0754); err != nil {
			return err
		}
	}

	switch typeCmd {
	case "service":
		return generateService(module)
	case "controller":
		return generateController(module)
	case "module":
		return generateModule(module)
	}

	return nil
}

func generateService(module *Module) error {
	serviceFile, err := os.Create(fmt.Sprintf("%s/%s_service.go", module.ModName, module.ModName))
	if err != nil {
		return err
	}
	defer serviceFile.Close()

	serviceTemplate := template.Must(template.New("service").Parse(string(tpl.ServiceTemplate())))
	err = serviceTemplate.Execute(serviceFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func generateController(module *Module) error {
	controllerFile, err := os.Create(fmt.Sprintf("%s/%s_controller.go", module.ModName, module.ModName))
	if err != nil {
		return err
	}
	defer controllerFile.Close()

	controllerTemplate := template.Must(template.New("controller").Parse(string(tpl.ControllerTemplate())))
	err = controllerTemplate.Execute(controllerFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func generateModule(module *Module) error {
	err := generateService(module)
	if err != nil {
		return err
	}

	err = generateController(module)
	if err != nil {
		return err
	}

	moduleFile, err := os.Create(fmt.Sprintf("%s/%s_module.go", module.ModName, module.ModName))
	if err != nil {
		return err
	}
	defer moduleFile.Close()

	moduleTemplate := template.Must(template.New("module").Parse(string(tpl.ModuleTemplate())))
	err = moduleTemplate.Execute(moduleFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func validateCmdName(source string) string {
	i := 0
	l := len(source)
	// The output is initialized on demand, then first dash or underscore
	// occurs.
	var output string

	for i < l {
		if source[i] == '-' || source[i] == '_' {
			if output == "" {
				output = source[:i]
			}

			// If it's last rune and it's dash or underscore,
			// don't add it output and break the loop.
			if i == l-1 {
				break
			}

			// If next character is dash or underscore,
			// just skip the current character.
			if source[i+1] == '-' || source[i+1] == '_' {
				i++
				continue
			}

			// If the current character is dash or underscore,
			// upper next letter and add to output.
			output += string(unicode.ToUpper(rune(source[i+1])))
			// We know, what source[i] is dash or underscore and source[i+1] is
			// uppered character, so make i = i+2.
			i += 2
			continue
		}

		// If the current character isn't dash or underscore,
		// just add it.
		if output != "" {
			output += string(source[i])
		}
		i++
	}

	if output == "" {
		return source // source is initially valid name.
	}
	return output
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
