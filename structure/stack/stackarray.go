// Stack Array
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stacklinkedlistwithlist.go, stack_test.go

package stack

var stackArray []interface{}

// stackPush push to first index of array
func stackPush(n interface{}) {
	stackArray = append([]interface{}{n}, stackArray...)
}

// stackLength return length of array
func stackLength() int {
	return len(stackArray)
}

// stackPeak return last input of array
func stackPeak() interface{} {
	return stackArray[0]
}

// stackEmpty check array is empty or not
func stackEmpty() bool {
	return len(stackArray) == 0
}

// stackPop return last input and remove it in array
func stackPop() interface{} {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
