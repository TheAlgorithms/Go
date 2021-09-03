// Dynamic Array
// description: A dynamic array is quite similar to a regular array, but its size is modifiable during program runtime.
// details: for more details check out those links below here:
// geeks for geeks article : https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/
// We can mention that Dynamic Array is like Slice for more detail about `Go` Slices check those articles :
//	https://blog.golang.org/slices-intro
// 	https://blog.golang.org/slices
// authors [Wesllhey Holanda](https://github.com/wesllhey), [Milad](https://github.com/miraddo)
// see dynamicarray.go, dynamicarray_test.go

package dynamicarray

// errors: used to handle CheckRangeFromIndex function with a reasonable error value
import (
	"errors"
)

var defaultCapacity = 10

// DynamicArray structure
// size: length of array
// capacity: the maximum length of the segment
// elementData: an array of any type of data with interface
type DynamicArray struct {
	size        int
	capacity    int
	elementData []interface{}
}

// Put function is change/update the value in array with the index and new value
func (da *DynamicArray) Put(index int, element interface{}) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.elementData[index] = element

	return nil
}

// Add function is add new element to our array
func (da *DynamicArray) Add(element interface{}) {
	if da.size == da.capacity {
		da.NewCapacity()
	}

	da.elementData[da.size] = element
	da.size++
}

// Remove function is remove an element with the index
func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.elementData[index:], da.elementData[index+1:da.size])
	da.elementData[da.size-1] = nil

	da.size--

	return nil
}

// Get function is return one element with the index of array
func (da *DynamicArray) Get(index int) (interface{}, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.elementData[index], nil
}

// IsEmpty function is check that the array has value or not
func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

// GetData function return all value of array
func (da *DynamicArray) GetData() []interface{} {
	return da.elementData[:da.size]
}

// CheckRangeFromIndex function it will check the range from the index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

// NewCapacity function increase the capacity
func (da *DynamicArray) NewCapacity() {
	if da.capacity == 0 {
		da.capacity = defaultCapacity
	} else {
		da.capacity = da.capacity << 1
	}

	newDataElement := make([]interface{}, da.capacity)

	copy(newDataElement, da.elementData)

	da.elementData = newDataElement
}
