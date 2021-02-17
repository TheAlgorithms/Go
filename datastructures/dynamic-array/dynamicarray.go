package dynamicarray

import (
	"errors"
)

var defaultCapacity = 10

// DynamicArray structure
type DynamicArray struct {
	size        int
	capacity    int
	elementData []interface{}
}

// Put function
func (da *DynamicArray) Put(index int, element interface{}) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.elementData[index] = element

	return nil
}

// Add function
func (da *DynamicArray) Add(element interface{}) {
	if da.size == da.capacity {
		da.NewCapacity()
	}

	da.elementData[da.size] = element
	da.size++
}

// Remove function
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

// Get function
func (da *DynamicArray) Get(index int) (interface{}, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.elementData[index], nil
}

// IsEmpty function
func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

// GetData function
func (da *DynamicArray) GetData() []interface{} {
	return da.elementData[:da.size]
}

// CheckRangeFromIndex function
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

// NewCapacity function
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

// func main() {
// 	numbers := dynamicArray{}
// 	fmt.Println(numbers.isEmpty())

// 	numbers.add(10)
// 	numbers.add(20)
// 	numbers.add(30)
// 	numbers.add(40)
// 	numbers.add(50)

// 	fmt.Println(numbers.isEmpty())

// 	fmt.Println(numbers.getData())

// 	numbers.remove(1)

// 	fmt.Println(numbers.getData())

// 	numberFound, _ := numbers.get(1)
// 	fmt.Println(numberFound)

// 	numbers.put(0, 100)
// 	fmt.Println(numbers.getData())
// }
