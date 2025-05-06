package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate_Makefile(t *testing.T) {
	buf := new(bytes.Buffer)

	// Set command output to buffer
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"make"})

	// Execute the root command
	err := rootCmd.Execute()
	assert.Nil(t, err)

	generatedFile := filepath.Join(".", "Makefile")
	_, err = os.Stat(generatedFile)
	assert.Nil(t, err)

	os.Remove(generatedFile)
}
