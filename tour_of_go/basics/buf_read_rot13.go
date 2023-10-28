package basics

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	chars := []byte("AZaz")

	known_char := true
	start_char := chars[0]

	for i := 0; i < len(p); i++ {
		if p[i] >= chars[2] && p[i] <= chars[3] {
			start_char = chars[2]
			known_char = true
		} else if p[i] >= chars[0] && p[i] <= chars[1] {
			start_char = chars[0]
			known_char = true
		} else {
			known_char = false
		}
		if known_char {
			tmp := p[i] - start_char + 13
			p[i] = start_char + (tmp % 26)
		}
	}
	return
}

func _main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!  ")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
