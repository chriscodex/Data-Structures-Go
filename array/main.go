package main

import "fmt"

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

/* Auxiliar Methods */
func (a *array) shiftIndexBack(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
}

func (a *array) shiftIndexForward(index int) {
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.length++
}

// Delete the first element
func (a *array) DeleteFirstElement() any {
	element := a.data[0]
	delete(a.data, 0)
	a.shiftIndexBack(0)
	a.length--
	return element
}

// Delete the last element, return the element eliminated
func (a *array) DeleteLastElement() any {
	element := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return element
}

// Delete element by index
func (a *array) Delete(index int) any {
	element := a.data[index]
	delete(a.data, index)
	a.shiftIndexBack(index)
	a.length--
	return element
}

// Add element at the beginning
func (a *array) AddBegin(element any) any {
	a.shiftIndexForward(0)
	a.data[0] = element
	return element
}

// Add element at any index
func (a *array) AddElement(element any, index int) any {
	a.shiftIndexForward(index)
	a.data[index] = element
	return element
}

func main() {
	myArray := Constructor()
	myArray.Append(0)
	myArray.Append(1)
	myArray.Append(2)
	myArray.DeleteLastElement()
	for i := 0; i < myArray.length; i++ {
		fmt.Println(myArray.Get(i))
	}
	myArray.Append(2)
	myArray.Append(3)
	myArray.Append(4)
	fmt.Println(myArray)
	myArray.AddElement(5, 2)
	fmt.Println(myArray)
}
