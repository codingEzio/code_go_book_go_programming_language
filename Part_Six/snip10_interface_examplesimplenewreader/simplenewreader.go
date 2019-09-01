package simplenewreader

import (
	"io"
)

type stringReader struct {
	str string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.str)
	r.str = r.str[n:]

	if len(r.str) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(str string) io.Reader {
	return &stringReader{str}
}
