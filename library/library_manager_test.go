package library

import "testing"

func TestNewLibraryManager(t *testing.T) {
	libraryPath := "test/library/path"
	dbPath := "test/db/path"
	lm := NewLibraryManager(libraryPath, dbPath)

	if lm.LibraryPath != libraryPath {
		t.Errorf("Expected LibraryPath to be %s, got %s", libraryPath, lm.LibraryPath)
	}
}

// Additional tests can be written for `AddBookToLibrary`, `SearchLibrary`, and `UpdateBookDetails`
