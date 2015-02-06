package main

import "github.com/dupoxy/go-tour-fr/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	// Infinite flux of A
	b[0] = 'A'

	// Number of bits read, error
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
