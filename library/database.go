package library

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

// SetPath sets a new path for the database connection
func (db *Database) SetPath(newPath string) error {
	// Close the existing database connection if open
	if db.DB != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}
	// Open a new database connection with the new path
	newDb, err := sql.Open("sqlite3", newPath)
	if err != nil {
		return err
	}
	db.DB = newDb
	return nil
}

func NewDatabase(dbPath string) *Database {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{DB: db}
}

func (db *Database) InitSchema() error {
	query := `
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        author TEXT,
        genre TEXT,
        publication_date TEXT,
        publisher TEXT,
        language TEXT,
        isbn TEXT,
        page_count INTEGER,
        read BOOLEAN,
        rating INTEGER,
        notes TEXT,
        cover_image_path TEXT
    );`
	_, err := db.Exec(query)
	return err
}

func (db *Database) AddBook(metadata Metadata) error {
	query := `INSERT INTO books (title, author, genre, publication_date, publisher, language, isbn, page_count, read, rating, notes, cover_image_path) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, metadata.Title, metadata.Author, metadata.Genre, metadata.PublicationDate, metadata.Publisher, metadata.Language, metadata.ISBN, metadata.PageCount, metadata.Read, metadata.Rating, metadata.Notes, metadata.CoverImagePath)
	return err
}

func (db *Database) SearchBooks(searchTerm string) ([]Metadata, error) {
	query := `SELECT title, author, genre, publication_date, publisher, language, isbn, page_count, read, rating, notes, cover_image_path FROM books WHERE title LIKE ? OR author LIKE ?`
	rows, err := db.Query(query, "%"+searchTerm+"%", "%"+searchTerm+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Metadata
	for rows.Next() {
		var metadata Metadata
		err = rows.Scan(&metadata.Title, &metadata.Author, &metadata.Genre, &metadata.PublicationDate, &metadata.Publisher, &metadata.Language, &metadata.ISBN, &metadata.PageCount, &metadata.Read, &metadata.Rating, &metadata.Notes, &metadata.CoverImagePath)
		if err != nil {
			return nil, err
		}
		books = append(books, metadata)
	}
	return books, nil
}
