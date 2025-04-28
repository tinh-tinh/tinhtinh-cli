package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	buf := new(bytes.Buffer)

	// Set command output to buffer
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"init"})

	// Execute the root command
	err := rootCmd.Execute()
	assert.Nil(t, err)

	_, err = os.Stat("main.go")
	assert.Nil(t, err)

	generatedFile := filepath.Join("app", "app_module.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	generatedFile = filepath.Join("app", "app_controller.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	generatedFile = filepath.Join("app", "app_service.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	os.RemoveAll("app")
	os.Remove("main.go")
}
