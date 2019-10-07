package deduplicator

import (
	"../supplier"
	"../storage"
	"fmt"
)

type Deduplicator struct {
	Data storage.Simple
	Supl supplier.Supplier
}

func (d *Deduplicator) Run() {
	dedupCache := make([]byte, 1024)
	memUsed := uint64(0)

	for _, err := d.Supl.Reader().Read(dedupCache); err == nil; {
		d.Data.Put(dedupCache)
		memUsed += 1024
		if memUsed % 100 == 0 {
			fmt.Println("mem used: ", memUsed/(1024*1024))
		}
	}
}
