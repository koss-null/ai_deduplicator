package supplier

import (
	"testing"
	"time"
)

func TestRingedSupplierExecutionTime(t *testing.T) {
	ringSizeBytes := 1024 * 1024 // 1MB
	spl, _ := NewRingedSupplyer(ringSizeBytes / 1024, nil)
	buffer := make([]byte, ringSizeBytes)
	start := time.Now()
	for i := 0; i < 1024; i++ { // reading 1GB from buffer 1kb size
		spl.Reader().Read(buffer)
	}
	t.Log(time.Now().Sub(start).Nanoseconds())

	spl, _ = NewRingedSupplyer(ringSizeBytes, nil)
	start = time.Now()
	for i := 0; i < 1024; i++ { // reading 1GB from buffer 1Mb size
		spl.Reader().Read(buffer)
	}
	t.Log(time.Now().Sub(start).Nanoseconds())

	buffer = make([]byte, 1)
	start = time.Now()
	for i := 0; i < 1024*1024*10; i++ { // reading 10MB from buffer 1Mb size
		spl.Reader().Read(buffer)
	}
	t.Log(time.Now().Sub(start).Nanoseconds())
}