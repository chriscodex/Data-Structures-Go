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

// Add an element at the last of the array
func (a *array) Append(item any) map[int]any {
	a.data[a.length] = item
	a.length++
	return a.data
}
