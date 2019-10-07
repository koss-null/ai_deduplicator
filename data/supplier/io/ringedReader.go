package io

import (
	"io"
)

type RingedReader struct {
	ring   *[]byte
	cursor *int
}

/*
	takes byte array returned from NewRingedWriter as a parameter
 */
func NewRingedReader(data *[]byte, nowOn *int) io.Reader {
	return &RingedReader{data, nowOn}
}

func (rr *RingedReader) Read(p []byte) (int, error) {
	if len(p) == 1 {
		p[0] = rr.get()
		return 1, nil
	}

	data := rr.getInterval(len(p))
	for i := range p {
		p[i] = data[i]
	}
	return len(p), nil
}

func (rr *RingedReader) get() byte {
	index := *rr.cursor % len(*rr.ring)
	*rr.cursor++
	return (*rr.ring)[index]
}

func (rr *RingedReader) getInterval(amount int) []byte {
	ret := make([]byte, 0, amount)

	start := *rr.cursor % len(*rr.ring)
	for amount != 0 {
		for amount != 0 && start < len(*rr.ring) {
			ret = append(ret, (*rr.ring)[start])
			start++
			amount--
		}
		if amount != 0 {
			start = 0
		}
	}

	*rr.cursor = start
	return ret
}