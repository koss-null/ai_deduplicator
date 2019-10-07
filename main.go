package main

import(
	"./data/deduplicator"
	"./data/storage"
	"./data/supplier"
)

func main() {
	//args := os.Args[1:]
	sp, _ := supplier.NewRingedSupplyer(2048, nil)
	dedup := deduplicator.Deduplicator{storage.NewSimpleStorage(), sp}
	dedup.Run()
}
