package library

import (
    "io"
    "os"
    "path/filepath"
)

// AddBook copies the e-book file to the library directory
func (lm *LibraryManager) AddBook(originalPath string) error {
    fileName := filepath.Base(originalPath)
    destPath := filepath.Join(lm.LibraryPath, fileName)

    // Open the source file
    srcFile, err := os.Open(originalPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    // Create the destination file
    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    // Copy the file
    _, err = io.Copy(destFile, srcFile)
    return err
}
