package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestChmod(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "filename.txt")
	f, err := os.Create(filePath)
	assert.NoError(t, err)
	assert.NoError(t, f.Chmod(0000))
	assert.NoError(t, f.Close())
	t.Cleanup(func() { assert.NoError(t, os.Remove(filePath)) })

	_, err = os.Open(filePath)
	assert.Error(t, err)
}