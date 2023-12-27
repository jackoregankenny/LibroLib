package library

import (
	"errors"
	"path/filepath"

	"github.com/bmaupin/go-epub"
	"github.com/unidoc/unipdf/v3/model"
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

// CheckFileType checks if the file is an e-book (PDF or EPUB)
func CheckFileType(filePath string) bool {
	extension := filepath.Ext(filePath)
	return extension == ".pdf" || extension == ".epub"
}

// ExtractMetadata determines the file type and extracts metadata
func (lm *LibraryManager) ExtractMetadata(filePath string) (Metadata, error) {
	if CheckFileType(filePath) == ".epub" {
		return ExtractMetadataFromEPUB(filePath)
	} else if CheckFileType(filePath) == ".pdf" {
		return ExtractMetadataFromPDF(filePath)
	}
	return Metadata{}, errors.New("unsupported file type")
}

// ExtractMetadataFromEPUB extracts metadata from an EPUB file
func ExtractMetadataFromEPUB(filePath string) (Metadata, error) {
	var metadata Metadata

	epub, err := epub.Open(filePath)
	if err != nil {
		return metadata, err
	}
	defer epub.Close()

	metadata.Title = epub.Title
	metadata.Author = epub.Creator
	metadata.PublicationDate = epub.Published
	metadata.Publisher = epub.Publisher
	metadata.Language = epub.Language
	// Other fields like Genre, ISBN, etc. might need manual input

	return metadata, nil
}

// ExtractMetadataFromPDF extracts metadata from a PDF file
func ExtractMetadataFromPDF(filePath string) (Metadata, error) {
	var metadata Metadata

	pdfDoc, err := model.NewPdfReaderFromFile(filePath, nil)
	if err != nil {
		return metadata, err
	}

	docInfo, err := pdfDoc.GetPdfProducer()
	if err != nil {
		return metadata, err
	}

	metadata.Title = docInfo.GetTitle()
	metadata.Author = docInfo.GetAuthor()
	// PDFs generally don't contain genre, rating, or read status.

	return metadata, nil
}
