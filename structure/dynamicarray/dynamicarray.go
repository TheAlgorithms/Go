// Package dynamicarray
// A dynamic array is quite similar to a regular array, but its Size is modifiable during program runtime,
// very similar to how a slice in Go works. The implementation is for educational purposes and explains
// how one might go about implementing their own version of slices.
//
// For more details check out those links below here:
// GeeksForGeeks article : https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/
// Go blog: https://blog.golang.org/slices-intro
// Go blog: https://blog.golang.org/slices
// authors [Wesllhey Holanda](https://github.com/wesllhey), [Milad](https://github.com/miraddo)
// see dynamicarray.go, dynamicarray_test.go
package dynamicarray

import (
	"errors"
)

var defaultCapacity = 10

// DynamicArray structure
type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []any
}

// Put function is change/update the value in array with the index and new value
func (da *DynamicArray) Put(index int, element any) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.ElementData[index] = element

	return nil
}

// Add function is add new element to our array
func (da *DynamicArray) Add(element any) {
	if da.Size == da.Capacity {
		da.NewCapacity()
	}

	da.ElementData[da.Size] = element
	da.Size++
}

// Remove function is remove an element with the index
func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.ElementData[index:], da.ElementData[index+1:da.Size])
	da.ElementData[da.Size-1] = nil

	da.Size--

	return nil
}

// Get function is return one element with the index of array
func (da *DynamicArray) Get(index int) (any, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.ElementData[index], nil
}

// IsEmpty function is check that the array has value or not
func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

// GetData function return all value of array
func (da *DynamicArray) GetData() []any {
	return da.ElementData[:da.Size]
}

// CheckRangeFromIndex function it will check the range from the index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.Size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

// NewCapacity function increase the Capacity
func (da *DynamicArray) NewCapacity() {
	if da.Capacity == 0 {
		da.Capacity = defaultCapacity
	} else {
		da.Capacity = da.Capacity << 1
	}

	newDataElement := make([]any, da.Capacity)

	copy(newDataElement, da.ElementData)

	da.ElementData = newDataElement
}
