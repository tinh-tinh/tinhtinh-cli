package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [package] [git_url]",
	Short: "Initialize a new project, Go module, and clone commands (optional)",
	Long: `Initializes a new project structure, creates a folder with the package name,
runs 'go mod init' with the specified package path, optionally clones a
repository into the created folder, and then executes 'go mod tidy'.
If no package path is provided, it attempts to infer it from the project
directory. If a Git URL is provided, it will be cloned into the package folder.`,
	Args: cobra.MaximumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) >= 2 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		var packageName string
		var gitURL string

		if len(args) > 0 {
			packageName = args[0]
		} else {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current working directory:", err)
				return
			}
			packageName = filepath.Base(wd)
			fmt.Printf("Using directory name '%s' as module path. You can specify a different path as an argument.\n", packageName)
		}

		if len(args) > 1 {
			gitURL = args[1]
		}

		projectDir := packageName // Use the package name as the project directory

		fmt.Printf("Creating project directory: %s\n", projectDir)
		err := os.MkdirAll(projectDir, 0755) // Create the directory with read/write/execute permissions for the owner and read/execute for others
		if err != nil {
			fmt.Printf("Error creating project directory '%s': %v\n", projectDir, err)
			return
		}

		// Change the current working directory to the newly created project directory
		err = os.Chdir(projectDir)
		if err != nil {
			fmt.Printf("Error changing directory to '%s': %v\n", projectDir, err)
			return
		}
		defer os.Chdir("..") // Change back to the original directory when the function exits

		fmt.Println("Initializing project...")
		_, err = initializeProject() // Assuming this function now works within the project directory
		if err != nil {
			fmt.Println("Error during project initialization:", err)
			return
		}

		fmt.Printf("Initializing Go module with path: %s\n", packageName)
		cmdGoModInit := exec.Command("go", "mod", "init", packageName)
		cmdGoModInit.Stdout = os.Stdout
		cmdGoModInit.Stderr = os.Stderr
		err = cmdGoModInit.Run()
		if err != nil {
			fmt.Println("Error running 'go mod init':", err)
			fmt.Println("Make sure Go is installed and in your PATH.")
			return
		}
		fmt.Println("Go module initialized.")

		if gitURL != "" {
			fmt.Printf("Cloning repository '%s' into '%s'...\n", gitURL, projectDir)
			cmdGitClone := exec.Command("git", "clone", gitURL, ".") // Clone into the current directory
			cmdGitClone.Stdout = os.Stdout
			cmdGitClone.Stderr = os.Stderr
			err = cmdGitClone.Run()
			if err != nil {
				fmt.Printf("Error cloning repository '%s': %v\n", gitURL, err)
				fmt.Println("Make sure Git is installed and in your PATH, and the URL is correct.")
				return
			}
			fmt.Println("Repository cloned successfully.")
		}

		fmt.Println("Tidying Go module dependencies...")
		cmdGoModTidy := exec.Command("go", "mod", "tidy")
		cmdGoModTidy.Stdout = os.Stdout
		cmdGoModTidy.Stderr = os.Stderr
		err = cmdGoModTidy.Run()
		if err != nil {
			fmt.Println("Error running 'go mod tidy':", err)
			return
		}
		fmt.Println("Go module dependencies tidied.")

		fmt.Printf("🚀 TinhTinh initialized successfully in '%s' with Go modules and dependencies tidied!", projectDir)
	},
}

func initializeProject() (string, error) {
	modName := getModImportPath()
	project := &Project{
		PkgName: modName,
	}

	if err := project.Create(); err != nil {
		return "", err
	}

	return modName, nil
}

func getModImportPath() string {
	mod, cd := parseModInfo()
	return path.Join(mod.Path, fileToURL(strings.TrimPrefix(cd.Dir, mod.Dir)))
}

func fileToURL(in string) string {
	i := strings.Split(in, string(filepath.Separator))
	return path.Join(i...)
}

func parseModInfo() (Mod, CurDir) {
	var mod Mod
	var dir CurDir

	m := modInfoJSON("-m")
	cobra.CheckErr(json.Unmarshal(m, &mod))

	if mod.Path == "command-line-arguments" {
		cobra.CheckErr("Please run `go mod init <MODNAME>` before `cobra-cli init`")
	}

	e := modInfoJSON("-e")
	cobra.CheckErr(json.Unmarshal(e, &dir))

	return mod, dir
}

type Mod struct {
	Path, Dir, GoMod string
}

type CurDir struct {
	Dir string
}

// func goGet(mod string) error {
// 	return exec.Command("go", "get", mod).Run()
// }

func modInfoJSON(args ...string) []byte {
	cmdArgs := append([]string{"list", "-json"}, args...)
	out, err := exec.Command("go", cmdArgs...).Output()
	cobra.CheckErr(err)

	return out
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
