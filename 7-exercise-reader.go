package main

import "github.com/dupoxy/go-tour-fr/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
    s = s[:cap(s)];
    for i := range s {
        s[i] = 'A'
    }
    return cap(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
