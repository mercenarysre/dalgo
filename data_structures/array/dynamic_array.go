package main

import (
	"fmt"
)

type DynamicArray struct {
	Capacity int // memory size allocated for array
	Length   int // number of real values in the array
	Arr      []int
}

// constructor function
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		// initial max no of elements the array can hold before resizing
		Capacity: 2,
		// Length indicating the array is empty initially
		Length: 0,
		// actual underlying slice that will store your data, with an initial capacity of 2
		Arr: make([]int, 2),
	}
}

// read/print all values in the array
func (d *DynamicArray) Print() {
	for i := 0; i < d.Length; i++ {
		fmt.Println(d.Arr[i])
	}
}

// reading a value in the array using an index
func (d *DynamicArray) Get(i int) int {
	if i < d.Length {
		return d.Arr[i]
	}
	// If i is greater than or equal to d.Length
	return -1
}

// search for a value in the array
func (d *DynamicArray) Search(n int) {
	for i := 0; i < d.Length; i++ {
		if d.Arr[i] == n {
			fmt.Println(i)
		} else {
			fmt.Println("The value does not exist in the array")
		}
	}
}

// Resizing the array
func (d *DynamicArray) Resize() {
	d.Capacity = d.Capacity * 2
	newArr := make([]int, d.Capacity)

	copy(newArr, d.Arr)
	d.Arr = newArr
}

func (d *DynamicArray) Pushback(n int) {
	if d.Length == d.Capacity {
		d.Resize()
	}
	// Length is less than Capacity normally
	// d.Length happens to be the index of the next available position in the array
	d.Arr[d.Length] = n
	// This updates the Length to reflect that there is now one more element in the dynamic array
	d.Length++
}

// insert value at the beginning where i = 0
func (d *DynamicArray) Insert(i, n int) []int {
	if d.Length == d.Capacity {
		d.Resize()
	}

	// Shift starting from the end to i
	for index := d.Length - 1; index > i-1; index-- {
		d.Arr[index+1] = d.Arr[index]
	}

	// Insert at i
	d.Arr[i] = n

	return d.Arr
}

// Deleting a value at the beginning where i = 0
func (d *DynamicArray) Remove(i int) []int {
	// Shift starting from i + 1 to end
	for index := i + 1; index < d.Length; index++ {
		d.Arr[index-1] = d.Arr[index]
	}

	return d.Arr[:d.Length-1]
}

// inserting a value in the array
func (d *DynamicArray) insert(i, n int) {
	if i < d.Length {
		d.Arr[i] = n
	}
}

// deleting a value from the end of the array
func (d *DynamicArray) Popback() {
	if d.Length > 0 {
		d.Length--
	}
}
