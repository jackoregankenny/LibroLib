package library

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAddBookFile(t *testing.T) {
	// Setup a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}
	defer os.RemoveAll(tempDir)

	lm := &LibraryManager{LibraryPath: tempDir}
	testFilePath := filepath.Join(tempDir, "test.epub")
	if _, err := os.Create(testFilePath); err != nil {
		t.Fatal("Failed to create test file:", err)
	}

	if err := lm.AddBook(testFilePath); err != nil {
		t.Errorf("AddBook() failed: %v", err)
	}

	// Further checks can be added to ensure the file was copied correctly
}

// Additional error case tests can be added
