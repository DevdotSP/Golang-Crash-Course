package diffuc

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ExampleFileDirectory demonstrates common file and directory operations
func ExampleFileDirectory() {
	// Uncomment the functions to test specific examples:
	// FirstExample()                  // Create an empty file
	// ToCreateDirectory()             // Create a new directory
	// ToRenameFile()                  // Rename a file
	// ToMoveToOtherFile()             // Move a file to another location
	// ToCopyFile()                    // Copy a file
	// ToReadDetails()                 // Read file metadata
	// ToRemoveFile()                  // Delete a file
	// ReadfileCharacter()             // Read file character by character
	// ToTruncate()                    // Truncate file to a specific size
	// ToAppendContentToTextFile()     // Append content to a file
	// ToChangeFilePermission()        // Change file permissions, ownership, timestamps
	// TestAppendZipFiles()            // Create a zip archive with multiple files
	// ReadZipFile()                   // List files in a zip archive
	ToUnzipFile() // Extract all files from a zip archive
}

// FirstExample creates an empty file
func FirstExample() {
	emptyFile, err := os.Create("Files/empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created file:", emptyFile.Name())
	emptyFile.Close()
}

// ToCreateDirectory creates a directory if it does not exist
func ToCreateDirectory() {
	_, err := os.Stat("Test")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("Test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
		log.Println("Directory 'Test' created successfully")
	}
}

// ToRenameFile renames a file from oldName to newName
func ToRenameFile() {
	oldName := "Test/test.txt"
	newName := "Test/testing.txt"

	fmt.Println("Renaming:", oldName, "=>", newName)
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File renamed successfully")
}

// ToMoveToOtherFile moves a file to a new location
func ToMoveToOtherFile() {
	oldLocation := "Test/testing.txt"
	newLocation := "Files/testing.txt"

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File moved successfully")
}

// ToCopyFile copies a file from source to destination
func ToCopyFile() {
	sourceFile, err := os.Open("Files/testing.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create("Test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes successfully.", bytesCopied)
}

// ToReadDetails prints metadata about a file
func ToReadDetails() {
	fileStat, err := os.Stat("Files/empty.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())
	fmt.Println("Size:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last Modified:", fileStat.ModTime())
	fmt.Println("Is Directory:", fileStat.IsDir())
}

// ToRemoveFile deletes a file
func ToRemoveFile() {
	err := os.Remove("Files/testing.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File deleted successfully")
}

// ReadfileCharacter reads a file character by character
func ReadfileCharacter() {
	filename := "Files/empty.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := bufio.NewScanner(strings.NewReader(string(filebuffer)))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
	fmt.Println()
}

// ToTruncate truncates a file to a specific length
func ToTruncate() {
	err := os.Truncate("Files/empty.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File truncated successfully")
}

// ToAppendContentToTextFile appends text content to a file
func ToAppendContentToTextFile() {
	message := "Add this content at end"
	filename := "Files/empty.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
	log.Println("Content appended successfully")
}

// ToChangeFilePermission changes file permissions, ownership, and timestamps
func ToChangeFilePermission() {
	_, err := os.Stat("Files/empty.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exists")

	// Change permissions
	err = os.Chmod("Files/empty.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change ownership (Linux)
	err = os.Chown("Files/empty.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change timestamps
	addOneDay := time.Now().Add(24 * time.Hour)
	err = os.Chtimes("Files/empty.txt", addOneDay, addOneDay)
	if err != nil {
		log.Println(err)
	}
}

// appendFiles adds a file to the zip archive
func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create zip entry for %s: %s", filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filename, err)
	}
	return nil
}

// TestAppendZipFiles creates a zip file with multiple files
func TestAppendZipFiles() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("Test/test.zip", flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	files := []string{"Test/test1.txt", "Test/test2.txt", "Test/test3.txt"}

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
	log.Println("Zip file created successfully")
}

// listFiles prints the name of a file in the zip archive
func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open zip file %s: %s", file.Name, err)
	}
	defer fileread.Close()

	fmt.Println("File in zip:", file.Name)
	return nil
}

// ReadZipFile lists all files inside a zip archive
func ReadZipFile() {
	read, err := zip.OpenReader("Test/test.zip")
	if err != nil {
		log.Fatalf("Failed to open zip: %s", err)
	}
	defer read.Close()

	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}

// ToUnzipFile extracts all files from a zip archive
func ToUnzipFile() {
	zipReader, err := zip.OpenReader("Test/test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		extractedFilePath := filepath.Join("./", file.Name)

		if file.FileInfo().IsDir() {
			log.Println("Directory created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
