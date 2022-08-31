package main

type HashTable struct {
	data [][][]string
}

// Constructor
func NewHashTable(size int) HashTable {
	return HashTable{
		data: make([][][]string, size),
	}
}
