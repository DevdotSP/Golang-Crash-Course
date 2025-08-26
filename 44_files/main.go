package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// ---------------------------------------------------------
	// (1) CREATE A FILE
	// ---------------------------------------------------------
	// WHAT: Creates a new file "a.txt".
	// WHY: Needed if the file doesn’t exist yet.
	// WHEN: Writing fresh logs, reports, or exports.
	// NOTE: If the file already exists, it will be truncated (emptied).
	newFile, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err) // idiomatic way to handle errors in Go
	}

	// ---------------------------------------------------------
	// (2) TRUNCATE A FILE
	// ---------------------------------------------------------
	// WHAT: Removes all content from "a.txt" (size becomes 0).
	// WHY: Useful if you want to keep the file but clear its contents.
	// WHEN: Resetting logs or starting fresh without deleting the file.
	err = os.Truncate("a.txt", 0) // 0 means completely empty the file
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------
	// (3) CLOSE A FILE
	// ---------------------------------------------------------
	// WHAT: Closes the file once you’re done with it.
	// WHY: Prevents memory leaks and file locking issues.
	// WHEN: Always close after reading/writing is finished.
	newFile.Close()

	// ---------------------------------------------------------
	// (4) OPEN AN EXISTING FILE (READ-ONLY)
	// ---------------------------------------------------------
	// WHAT: Opens "a.txt" in read-only mode.
	// WHY: Safest way to read without modifying the file.
	// WHEN: When you only want to parse or inspect contents.
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// ---------------------------------------------------------
	// (5) OPEN A FILE WITH CUSTOM OPTIONS
	// ---------------------------------------------------------
	// WHAT: Use os.OpenFile with flags for more control.
	// WHY: Lets you choose read/write, append, create, etc.
	// WHEN: Flexible situations like log files or updating data.
	//
	// Common flags:
	//   os.O_RDONLY  → Read only
	//   os.O_WRONLY  → Write only
	//   os.O_RDWR    → Read + Write
	//   os.O_APPEND  → Append to end
	//   os.O_CREATE  → Create if not exist
	//   os.O_TRUNC   → Clear file on open
	//
	// Example: os.O_CREATE|os.O_APPEND
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// ---------------------------------------------------------
	// (6) GET FILE INFORMATION
	// ---------------------------------------------------------
	// WHAT: Retrieve file metadata like size, permissions, etc.
	// WHY: Useful to check before processing a file.
	// WHEN: Logging file stats or validating before reads/writes.
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => a.txt
	p("Size in bytes:", fileInfo.Size())    // => 0
	p("Last modified:", fileInfo.ModTime()) // timestamp
	p("Is Directory? ", fileInfo.IsDir())   // false
	p("Permissions:", fileInfo.Mode())      // e.g. -rw-r-----

	// ---------------------------------------------------------
	// (7) CHECK IF A FILE EXISTS
	// ---------------------------------------------------------
	// WHAT: Try to Stat() "b.txt" and handle errors.
	// WHY: Prevents crashing when accessing missing files.
	// WHEN: Safe way to check existence before opening.
	_, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("The file does not exist")
		}
	}

	// ---------------------------------------------------------
	// (8) RENAME OR MOVE A FILE
	// ---------------------------------------------------------
	// WHAT: Renames "a.txt" → "aaa.txt".
	// WHY: Useful for versioning, organizing, or moving.
	// WHEN: Archiving or rotating files.
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------
	// (9) REMOVE A FILE
	// ---------------------------------------------------------
	// WHAT: Permanently deletes a file from the filesystem.
	// WHY: Cleanup temporary files or outdated data.
	// WHEN: Always double-check before deleting!
	err = os.Remove("aa.txt")
	if err != nil {
		log.Fatal(err)
	}
}
