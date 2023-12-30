package library

import "testing"

func TestCheckFileType(t *testing.T) {
	epubFile := "book.epub"
	nonEpubFile := "image.png"

	if !CheckFileType(epubFile) {
		t.Errorf("CheckFileType() failed for EPUB file: %s", epubFile)
	}

	if CheckFileType(nonEpubFile) {
		t.Errorf("CheckFileType() incorrectly identified non-EPUB file: %s", nonEpubFile)
	}
}

// Additional tests can be written for `ExtractMetadata` and `ExtractMetadataFromEPUB`
