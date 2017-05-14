package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {
	count := 0
	for i := range b {
		b[i] = 'A'
		count = i
	}
	return count, nil
}

func main() {
	reader.Validate(MyReader{})
}
