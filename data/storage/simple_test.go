package storage

import (
	"testing"
	"math"
	"fmt"
)

func TestSimpleStorage(t *testing.T) {
	s := NewSimpleStorage()

	cashGB := make([]byte, 1024*1024*1024)
	for i := range cashGB {
		cashGB[i] = byte(i % math.MaxInt8)
	}

	if err := s.Put(cashGB); err != nil {
		t.Error(err.Error())
	}

	result := s.Get(0, 1024*1024*1024)
	for i := range result {
		for j := range result[i] {
			t.Log(result[i])
			if result[i][j] != cashGB[i * len(result[i]) + j] {
				t.Error(fmt.Sprintf("data corruted after storing: expected: %d, got: %d\n",
					cashGB[i * len(result[i]) + j], result[i][j]))
			}
		}
	}
}
