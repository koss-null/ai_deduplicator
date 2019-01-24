package writer

import (
	"io"
	"math/rand"
	"encoding/binary"
)

type RingedWriter struct {
	ring []byte
}

func (rw RingedWriter) Write(p []byte) (n int, err error) {
	return len(rw.ring), nil
}

/*
	fillBytes() -- fills the ring block with the data, generated by gedData function
	in case this function stops returning any data, but the ring[] is not filled,
	all other data fills with zeroes
 */
func fillBytes(ringSize int, getData func() []byte) []byte {
	ring := make([]byte, ringSize, ringSize)
	curLength := 0
	for curLength < ringSize {
		data := getData()
		if len(data) == 0 {
			for i := curLength; i < ringSize; i++ {
				ring[i] = 0
			}
			curLength = ringSize;
		}
		for i := curLength; i < ringSize && i-curLength < len(data); i++ {
			ring[i] = data[i-curLength]
		}
		curLength += len(data)
	}

	return ring
}

func NewRingedWriter(ringSize int, source io.Reader) io.Writer {
	var getData func() []byte
	if source != nil {
		// read from source ring
		getData = func() []byte {
			data := make([]byte, ringSize)
			n, err := source.Read(data)
			if err != nil {
				for i := range data {
					data[i] = 0
				}
				return data
			}
			if n < ringSize {
				for i := n; i < ringSize; i++ {
					data[i] = 0
				}
				return data
			}
			return data[0:ringSize]
		}
	} else {
		// random ring
		getData = func() []byte {
			data := make([]byte, ringSize)
			for i := 0; i < ringSize; {
				bs := make([]byte, 8)
				binary.LittleEndian.PutUint64(bs, rand.Uint64())
				for j := 0; j < len(bs); j++ {
					data[i + j] = bs[j]
				}
				i += len(bs)
			}
			return data
		}
	}

	return RingedWriter{fillBytes(ringSize, getData)}
}