package supplier

import (
	"io"
	"errors"
	customIO "./io"
)

type (
	Supplier interface {
		Reader() io.Reader
		Writer() io.Writer
	}

	supplier struct {
		reader io.Reader
		writer io.Writer
	}
)

func NewRingedSupplyer(ringSize int, source io.Reader) (Supplier, error) {
	if ringSize < 1 {
		return nil, errors.New("ring size is too small (should be at least 1)")
	}
	writer, dataLink, cursorLink := customIO.NewRingedWriter(ringSize, source)
	reader := customIO.NewRingedReader(dataLink, cursorLink)

	return &supplier{reader, writer}, nil
}

func (s *supplier) Reader() io.Reader {
	return s.reader
}

func (s *supplier) Writer() io.Writer {
	return s.writer
}