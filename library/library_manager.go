package library

import (
	"errors"
	"os"
	"path/filepath"
)

type LibraryManager struct {
	LibraryPath string
	db          *Database
}

// NewLibraryManager creates and initializes a new LibraryManager
func NewLibraryManager(libraryPath, dbPath string) *LibraryManager {
	var db *Database
	if dbPath != "" {
		db = NewDatabase(dbPath)
		db.InitSchema() // Initialize database schema
	}

	return &LibraryManager{
		LibraryPath: libraryPath,
		db:          db,
	}
}

// AddBookToLibrary adds a new book to the library
func (lm *LibraryManager) AddBookToLibrary(filePath string) error {
	// Check if the file type is supported
	if !CheckFileType(filePath) {
		return errors.New("unsupported file type")
	}

	// Add book file to library folder
	if err := lm.AddBook(filePath); err != nil {
		return err
	}

	// Extract metadata
	metadata, err := lm.ExtractMetadata(filePath)
	if err != nil {
		return err
	}

	// Add metadata to database
	return lm.db.AddBook(metadata)
}

// SearchLibrary searches for books in the library
func (lm *LibraryManager) SearchLibrary(searchTerm string) ([]Metadata, error) {
	return lm.db.SearchBooks(searchTerm)
}

// Additional functions for organizing, updating, or retrieving book details  here

// UpdateBookDetails updates the metadata of an existing book in the library
func (lm *LibraryManager) UpdateBookDetails(bookID int, updatedMetadata Metadata) error {
	query := `UPDATE books SET title=?, author=?, genre=?, publication_date=?, publisher=?, language=?, isbn=?, page_count=?, read=?, rating=?, notes=?, cover_image_path=? WHERE id=?`
	_, err := lm.db.Exec(query, updatedMetadata.Title, updatedMetadata.Author, updatedMetadata.Genre, updatedMetadata.PublicationDate, updatedMetadata.Publisher, updatedMetadata.Language, updatedMetadata.ISBN, updatedMetadata.PageCount, updatedMetadata.Read, updatedMetadata.Rating, updatedMetadata.Notes, updatedMetadata.CoverImagePath, bookID)
	return err
}

// DeleteBook removes a book from the library
func (lm *LibraryManager) DeleteBook(bookID int) error {
	// First, retrieve the file name from the database
	var fileName string
	query := `SELECT cover_image_path FROM books WHERE id = ?`
	row := lm.db.QueryRow(query, bookID)
	err := row.Scan(&fileName)
	if err != nil {
		return err
	}

	// Delete the book file from the library folder
	filePath := filepath.Join(lm.LibraryPath, fileName)
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	// Then, remove the book entry from the database
	query = `DELETE FROM books WHERE id = ?`
	_, err = lm.db.Exec(query, bookID)
	return err
}

// OrganizeLibrary organizes the library files in a specified manner
func (lm *LibraryManager) OrganizeLibrary(organizationMethod string) error {
	// Example implementation: organize by author
	if organizationMethod == "author" {
		rows, err := lm.db.Query("SELECT id, author, cover_image_path FROM books")
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var author, filePath string
			err = rows.Scan(&id, &author, &filePath)
			if err != nil {
				return err
			}

			authorDir := filepath.Join(lm.LibraryPath, author)
			if err := os.MkdirAll(authorDir, 0755); err != nil {
				return err
			}

			newFilePath := filepath.Join(authorDir, filepath.Base(filePath))
			if err := os.Rename(filepath.Join(lm.LibraryPath, filePath), newFilePath); err != nil {
				return err
			}

			// Update the new file path in the database
			_, err = lm.db.Exec("UPDATE books SET cover_image_path = ? WHERE id = ?", newFilePath, id)
			if err != nil {
				return err
			}
		}
	} else if organizationMethod == "genre" {
		rows, err := lm.db.Query("SELECT id, genre, cover_image_path FROM books")
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var genre, filePath string
			err = rows.Scan(&id, &genre, &filePath)
			if err != nil {
				return err
			}

			genreDir := filepath.Join(lm.LibraryPath, genre)
			if err := os.MkdirAll(genreDir, 0755); err != nil {
				return err
			}

			newFilePath := filepath.Join(genreDir, filepath.Base(filePath))
			if err := os.Rename(filepath.Join(lm.LibraryPath, filePath), newFilePath); err != nil {
				return err
			}

			// Update the new file path in the database
			_, err = lm.db.Exec("UPDATE books SET cover_image_path = ? WHERE id = ?", newFilePath, id)
			if err != nil {
				return err
			}
		}
	} else if organizationMethod == "title" {
		rows, err := lm.db.Query("SELECT id, title, cover_image_path FROM books")
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var title, filePath string
			err = rows.Scan(&id, &title, &filePath)
			if err != nil {
				return err
			}

			titleDir := filepath.Join(lm.LibraryPath, title)
			if err := os.MkdirAll(titleDir, 0755); err != nil {
				return err
			}

			newFilePath := filepath.Join(titleDir, filepath.Base(filePath))
			if err := os.Rename(filepath.Join(lm.LibraryPath, filePath), newFilePath); err != nil {
				return err
			}

			// Update the new file path in the database
			_, err = lm.db.Exec("UPDATE books SET cover_image_path = ? WHERE id = ?", newFilePath, id)
			if err != nil {
				return err
			}
		}
	}

	func (lm *LibraryManager) GetBooks() ([]Metadata, error) {
		// Logic to retrieve and return all books from the database
		rows, err := lm.db.Query("SELECT id, title, author, genre, cover_image_path FROM books")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var books []Metadata
		for rows.Next() {
			var book Metadata
			err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.CoverImagePath)
			if err != nil {
				return nil, err
			}
			books = append(books, book)
		}

		return books, nil
	}

	return nil
}
