package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	buf := new(bytes.Buffer)

	// Set command output to buffer
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"generate", "module", "users"})

	// Execute the root command
	err := rootCmd.Execute()
	assert.Nil(t, err)

	tmpDir := "users"

	generatedFile := filepath.Join(tmpDir, "users_module.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	generatedFile = filepath.Join(tmpDir, "users_controller.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	generatedFile = filepath.Join(tmpDir, "users_service.go")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	os.RemoveAll(tmpDir)
	os.Remove("main.go")
}
