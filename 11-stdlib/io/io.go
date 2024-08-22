package io

import (
	"io"
)


func CountLetters(r io.Reader) (map[string]int , error) {
	buf := make([]byte, 2048) // a shared fixed size memory that allow you to read even from large values
	out := map[string]int {}

	for {
		n, err := r.Read(buf) // n allows us to know how much was read from file
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)] ++
			}
		}
		if err == io.EOF {
			// this is how you can tell you've read to the end
			// that's why we don't treat this like an error. 
			// C stuff
			return out, nil
		}
		if err != nil {
			return out, err
		}
	}

}

