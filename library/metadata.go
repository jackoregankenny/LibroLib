package library

import (
	"errors"
	"path/filepath"

	"github.com/taylorskalyo/goreader/epub"
)

// Metadata struct holds e-book metadata
type Metadata struct {
	Title           string
	Author          string
	Genre           string
	PublicationDate string // Format: YYYY-MM-DD
	Publisher       string
	Language        string
	ISBN            string
	PageCount       int
	Read            bool
	Rating          int // Rating on a scale of 1-5
	Notes           string
	CoverImagePath  string // File path or URL to the cover image
}

// CheckFileType checks if the file is an e-book (EPUB)
func CheckFileType(filePath string) bool {
	extension := filepath.Ext(filePath)
	return extension == ".epub"
}

// ExtractMetadata determines the file type and extracts metadata
func (lm *LibraryManager) ExtractMetadata(filePath string) (Metadata, error) {
	fileType := filepath.Ext(filePath)
	if fileType == ".epub" {
		return ExtractMetadataFromEPUB(filePath)
	}
	return Metadata{}, errors.New("unsupported file type")
}

func ExtractMetadataFromEPUB(filePath string) (Metadata, error) {
	var metadata Metadata

	rc, err := epub.OpenReader(filePath)
	if err != nil {
		return metadata, err
	}
	defer rc.Close()

	// Assuming there's only one rootfile (common case)
	book := rc.Rootfiles[0]

	// Extract metadata from the book
	metadata.Title = book.Title
	metadata.Author = book.Creator      // Assuming 'Creator' field exists for Author
	metadata.Language = book.Language   // Assuming 'Language' field exists
	metadata.Publisher = book.Publisher // Assuming 'Publisher' field exists
	// Add other necessary fields based on available goreader metadata fields.
	// For example, metadata.Author = book.Creator
	// Note: Adjust according to the actual fields available in the EPUB book struct

	return metadata, nil
}
