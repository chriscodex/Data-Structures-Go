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

// Get Method
func (h *HashTable) get(key string) string {
	address := h.hashMethod(key)
	currentBucket := h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				return currentBucket[i][1]
			}
		}
	}
	return ""
}

// Auxiliar Method
func remove(s [][]string, i int) [][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Delete Method
func (h *HashTable) delete(key string) {
	address := h.hashMethod(key)
	currentBucket := h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				h.data[address] = remove(currentBucket, i)
			}
		}
	}
}
