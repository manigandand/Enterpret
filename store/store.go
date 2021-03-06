package store

import (
	"enterpret/store/adaptee/inmemory"
	"enterpret/store/adapter"
	"log"
)

// Store global store connection interface
var Store adapter.Store

// Init loads the sample data and prepares the store layer
func Init() {
	// store inmemory adapter ...
	Store = inmemory.NewAdapter()
	if Store == nil {
		log.Fatalf("๐ฆ store initialize failed ๐")
	}
	log.Println("Inited Store...๐")
}
