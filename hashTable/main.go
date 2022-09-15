package main

import (
	"fmt"
)

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
func Remove(s [][]string, i int) [][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Delete Method
func (h *HashTable) Delete(key string) {
	address := h.hashMethod(key)
	currentBucket := h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				h.data[address] = Remove(currentBucket, i)
			}
		}
	}
}

// Get All Keys
func (h *HashTable) getAllKeys() []string {
	keys := []string{}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != nil {
			if len(h.data[i]) == 1 {
				keys = append(keys, h.data[i][0][0])
			} else {
				for j := 0; j < len(h.data[i]); j++ {
					keys = append(keys, h.data[i][j][0])
				}
			}
		}
	}
	return keys
}

func main() {
	myHashTable := NewHashTable(5)
	fmt.Println(myHashTable)
	myHashTable.set("Christian", "1998")
	myHashTable.set("Heber", "1961")
	myHashTable.set("Leo", "1994")
	myHashTable.set("Gian", "1931")
	myHashTable.set("Rodolfo", "1991")
	fmt.Println(myHashTable)
	t := myHashTable.get("Christian")
	fmt.Println(t)
	fmt.Println(myHashTable)
	ak := myHashTable.getAllKeys()
	fmt.Println(ak)
}
