package io

type SliceReader struct {
	data []byte
}

func (sr *SliceReader) Read(data []byte) (int, error) {
	for i := range sr.data {
		data[i] = sr.data[i]
	}
	return len(data), nil
}