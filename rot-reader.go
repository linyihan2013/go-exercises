package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, error := rot.r.Read(b)
	for i, v := range b {
		if v >= 'a' && v <= 'z' {
			v -= 13
			if v < 'a' {
				v = 26 + v
			}
		} else if v >= 'A' && v <= 'Z' {
			v += 13
			if v < 'A' {
				v = 26 + v
			}
		}
		b[i] = v
	}

	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
