package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// Open (or create if not exists) a file in write-only mode.
	// Flags:
	// - os.O_WRONLY → open for writing only
	// - os.O_CREATE → create the file if it doesn’t exist
	// Permissions (0644):
	// - Owner: read & write
	// - Group: read
	// - Others: read
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the file is closed at the end of main()
	defer file.Close()

	// Wrap the file with a buffered writer.
	// WHY: Writing directly to disk can be slow. 
	// Buffered writer accumulates writes in memory and flushes them in batches.
	bufferedWriter := bufio.NewWriter(file)

	// Prepare some data (a byte slice). 
	// These values are ASCII codes: 97 = 'a', 98 = 'b', 99 = 'c'.
	bs := []byte{97, 98, 99}

	// Write the byte slice to the buffer (NOT immediately to the file).
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not yet file): %d\n", bytesWritten)
	// Example output: 3

	// Check how much free space remains in the buffer.
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
	// Default buffer size is 4096, minus what we just wrote.

	// Write a string (instead of a byte slice) to the buffer.
	_, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	// Check how many bytes are currently buffered (waiting to be flushed).
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered (pending flush): %d\n", unflushedBufferSize)
	// Example: 24 → (3 from bs + 21 from string)

	// IMPORTANT: At this point, nothing is on disk yet.
	// Data is still in memory, inside the buffer.

	// Flush forces all buffered data to be written to the file.
	err = bufferedWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
