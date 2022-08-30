package main

/* Array structure */
type array struct {
	length int
	data   map[int]any
}

// Array constructur
func Constructor() array {
	return array{
		length: 0,
		data:   map[int]any{},
	}
}

// Get value by index
func (a *array) Get(index int) any {
	return a.data[index]
}
