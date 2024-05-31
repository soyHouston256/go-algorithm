package myArray

import "errors"

type MyArray struct {
	length int
	data   map[int]interface{}
}

func NewArray() *MyArray {
	return &MyArray{
		length: 0,
		data:   make(map[int]interface{}),
	}
}

func (a *MyArray) Get(index int) (interface{}, error) {
	val, ok := a.data[index]
	if !ok {
		return nil, errors.New("index out of range")
	}
	return val, nil
}
func (a *MyArray) Push(value interface{}) int {
	a.data[a.length] = value
	a.length++
	return a.length
}

func (a *MyArray) Pop() (interface{}, error) {
	if a.length == 0 {
		return nil, errors.New("empty array")
	}
	lastElement := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return lastElement, nil
}

func (a *MyArray) Delete(index int) (interface{}, error) {
	val, ok := a.data[index]
	if !ok {
		return nil, errors.New("index out of range")
	}
	a.ShiftItems(index)
	return val, nil
}

func (a *MyArray) ShiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
}
