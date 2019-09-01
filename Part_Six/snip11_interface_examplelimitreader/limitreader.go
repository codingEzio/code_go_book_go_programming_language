package limitreader

import (
	"io"
)

type limitReader struct {
	reader   io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p[:r.limit])

	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}

	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{reader: r, limit: limit}
}
