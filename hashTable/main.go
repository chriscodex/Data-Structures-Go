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

// Hash Function
func (h *HashTable) hashMethod(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % len(h.data)
	}
	return hash
}

// Set Method
func (h *HashTable) set(key string, value string) [][][]string {
	address := h.hashMethod(key)
	h.data[address] = append(h.data[address], []string{key, value})
	return h.data
}
