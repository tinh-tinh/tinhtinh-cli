package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var makefileCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate a Makefile",
	Run: func(cmd *cobra.Command, args []string) {
		content := `
BINARY_NAME=main

build:
	go build -o bin/$(BINARY_NAME) .

run:
	go run main.go

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/

dev:
	air
`
		err := os.WriteFile("Makefile", []byte(content), 0644)
		if err != nil {
			fmt.Println("Failed to write Makefile:", err)
			return
		}
		fmt.Println("Makefile generated.")
	},
}

func init() {
	rootCmd.AddCommand(makefileCmd)
}
