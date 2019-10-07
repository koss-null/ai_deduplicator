package deduplicator

import (
	"../supplier"
	"../storage"
	"fmt"
	"crypto/md5"
)

type Deduplicator struct {
	Data storage.Simple
	Supl supplier.Supplier
	hashMap map[string]int // int shows page with this hash code
}

func NewDeduplicator(data storage.Simple, supplier supplier.Supplier) Deduplicator {
	return Deduplicator{data, supplier, map[string]int{}}
}

func (d *Deduplicator) Dedup(origData []byte) []byte {
	fmt.Println(origData)
	md := md5.Sum(origData)
	_, ok := d.hashMap[fmt.Sprintf("%x", md)]
	if ok {
		return []byte{}
	} else {
		d.hashMap[fmt.Sprintf("%x", md)] = 1 // fixme
	}
	return origData
}

func (d *Deduplicator) Run() {
	dedupCache := make([]byte, 2048)
	memUsed := uint64(0)

	for _, err := d.Supl.Reader().Read(dedupCache); err == nil; {
		d.Data.Put(d.Dedup(dedupCache))
		memUsed += 1024
		if memUsed % 100 == 0 {
			fmt.Println("mem used: ", memUsed/(1024*1024))
		}
	}
}
