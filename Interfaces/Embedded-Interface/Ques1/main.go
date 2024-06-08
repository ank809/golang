package main

import "fmt"

// Define the Reader interface
type Reader interface {
	Read() string
}

// Define the Writer interface
type Writer interface {
	Write(string)
}

// Define the File struct
type File struct {
	content string
}

// Implement the Read method for the Reader interface
func (f *File) Read() string {
	return f.content
}

// Implement the Write method for the Writer interface
func (f *File) Write(data string) {
	f.content = data
}

// Define the ReaderWriter interface that embeds both Reader and Writer interfaces
type ReaderWriter interface {
	Reader
	Writer
}

// Function that uses both the Read and Write methods
func ReadWrite(rw ReaderWriter) string {
	rw.Write("Hello, Go!")
	return rw.Read()
}

func main() {
	// Create an instance of File
	var f File

	// Call ReadWrite with a pointer to f (since the methods have pointer receivers)
	content := ReadWrite(&f)

	// Print the content read from the file
	fmt.Println(content) // Output: Hello, Go!
}
