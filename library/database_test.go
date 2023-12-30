package library

import (
	"database/sql"
	"testing"
)

func TestAddBook(t *testing.T) {
	// This requires setting up a mock or in-memory database.
	// Example with an in-memory SQLite database.
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in-memory database:", err)
	}
	defer db.Close()

	database := &Database{DB: db}
	// Initialize schema
	if err := database.InitSchema(); err != nil {
		t.Fatal("Failed to initialize schema:", err)
	}

	// Add a book and check for errors
	metadata := Metadata{Title: "Test Book", Author: "Author"}
	if err := database.AddBook(metadata); err != nil {
		t.Errorf("AddBook() failed: %v", err)
	}

	// Further validation can be added to ensure the book was added correctly
}

// Additional tests can be written for `SearchBooks`
