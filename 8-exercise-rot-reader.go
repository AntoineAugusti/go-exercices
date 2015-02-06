package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rot13(b byte) byte {

	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		b += 13
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		b -= 13
	}

	return b
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	for i := 0; i < len(p); i++ {
		p[i] = Rot13(p[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
