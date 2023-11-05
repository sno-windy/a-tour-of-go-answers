package main

import (
	"io"
	"os"
	"strings"
)

func Rot13(char byte) byte {
	if char >= 'a' && char <= 'm' {
		return char + 13
	} else if char >= 'n' && char <= 'z' {
		return char - 13
	} else if char >= 'A' && char <= 'M' {
		return char + 13
	} else if char >= 'N' && char <= 'Z' {
		return char - 13
	} else {
		return char
	}
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
	 	return n, err
	 }
	rawStrings := b[:n]
	for i, v := range rawStrings {
	 	b[i] = Rot13(v)
	}
	return n, err
}	

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
