package supplier

import "io"

type (
	Supplier interface {
		Reader() io.Reader
		Writer(io.Writer)
	}

	supplier struct {
		writer io.Writer
	}
)

func (s supplier) Reader() io.Reader {}

func (s supplier) Writer(writer io.Writer) {
	s.writer = writer
}