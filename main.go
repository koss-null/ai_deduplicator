package main

import(
	"./data/deduplicator"
	"./data/storage"
	"./data/supplier"
)

func main() {
	//args := os.Args[1:]
	sp, _ := supplier.NewRingedSupplyer(7212, nil)
	dedup := deduplicator.NewDeduplicator(storage.NewSimpleStorage(), sp)
	dedup.Run()
}
