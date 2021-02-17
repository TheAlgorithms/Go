package dynamicarray

import (
	"errors"
)

var defaultCapacity = 10

type dynamicArray struct {
	size        int
	capacity    int
	elementData []interface{}
}

func (da *dynamicArray) put(index int, element interface{}) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.elementData[index] = element

	return nil
}

func (da *dynamicArray) add(element interface{}) {
	if da.size == da.capacity {
		da.newCapacity()
	}

	da.elementData[da.size] = element
	da.size++
}

func (da *dynamicArray) remove(index int) error {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.elementData[index:], da.elementData[index+1:da.size])
	da.elementData[da.size-1] = nil

	da.size--

	return nil
}

func (da *dynamicArray) get(index int) (interface{}, error) {
	err := da.checkRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.elementData[index], nil
}

func (da *dynamicArray) isEmpty() bool {
	return da.size == 0
}

func (da *dynamicArray) getData() []interface{} {
	return da.elementData[:da.size]
}

func (da *dynamicArray) checkRangeFromIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func (da *dynamicArray) newCapacity() {
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
