// Tests for the deque package.

package deque

import (
	"testing"
)

// Zero values
var zeroInt int
var zeroString string

type QueryStructure[T any] struct {
	queryType      string
	parameter      T
	expectedResult interface{}
	expectedError  error
}

type TestCaseData[T any] struct {
	description string
	queries     []QueryStructure[T]
}

func TestDeque(t *testing.T) {

	// Test cases with ints as params
	testCasesInt := []TestCaseData[int]{
		{
			description: "Test empty deque",
			queries: []QueryStructure[int]{
				{
					queryType:      "IsEmpty",
					parameter:      zeroInt,
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					parameter:      zeroInt,
					expectedResult: zeroInt,
					expectedError:  ErrEmptyDeQueue,
				},
				{
					queryType:      "Rear",
					parameter:      zeroInt,
					expectedResult: zeroInt,
					expectedError:  ErrEmptyDeQueue,
				},
				{
					queryType:      "DeQueueFront",
					parameter:      zeroInt,
					expectedResult: zeroInt,
					expectedError:  ErrEmptyDeQueue,
				},
				{
					queryType:      "DeQueueRear",
					parameter:      zeroInt,
					expectedResult: zeroInt,
					expectedError:  ErrEmptyDeQueue,
				},
				{
					queryType:      "Length",
					parameter:      zeroInt,
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test deque with one element",
			queries: []QueryStructure[int]{
				{
					queryType:      "EnQueueFront",
					parameter:      1,
					expectedResult: zeroInt,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroInt,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					parameter:      zeroInt,
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					parameter:      zeroInt,
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroInt,
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					parameter:      zeroInt,
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroInt,
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroInt,
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test deque with multiple elements",
			queries: []QueryStructure[int]{
				{
					queryType:      "EnQueueFront",
					parameter:      1,
					expectedResult: zeroInt,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueFront",
					parameter:      2,
					expectedResult: zeroInt,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueRear",
					parameter:      3,
					expectedResult: zeroInt,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueRear",
					parameter:      4,
					expectedResult: zeroInt,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroInt,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					parameter:      zeroInt,
					expectedResult: 2,
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					parameter:      zeroInt,
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroInt,
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					parameter:      zeroInt,
					expectedResult: 2,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueRear",
					parameter:      zeroInt,
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroInt,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroInt,
					expectedResult: 2,
					expectedError:  nil,
				},
			},
		},
	}

	// Test cases with strings as params
	testCasesString := []TestCaseData[string]{
		{
			description: "Test one element deque",
			queries: []QueryStructure[string]{
				{
					queryType:      "EnQueueFront",
					parameter:      "a",
					expectedResult: zeroString,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroString,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					parameter:      zeroString,
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					parameter:      zeroString,
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroString,
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					parameter:      zeroString,
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroString,
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroString,
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test multiple elements deque",
			queries: []QueryStructure[string]{
				{
					queryType:      "EnQueueFront",
					parameter:      "a",
					expectedResult: zeroString,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueFront",
					parameter:      "b",
					expectedResult: zeroString,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueRear",
					parameter:      "c",
					expectedResult: zeroString,
					expectedError:  nil,
				},
				{
					queryType:      "EnQueueRear",
					parameter:      "d",
					expectedResult: zeroString,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroString,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					parameter:      zeroString,
					expectedResult: "b",
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					parameter:      zeroString,
					expectedResult: "d",
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroString,
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					parameter:      zeroString,
					expectedResult: "b",
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueRear",
					parameter:      zeroString,
					expectedResult: "d",
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					parameter:      zeroString,
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					parameter:      zeroString,
					expectedResult: 2,
					expectedError:  nil,
				},
			},
		},
	}

	// Run tests with ints
	for _, testCase := range testCasesInt {
		t.Run(testCase.description, func(t *testing.T) {
			dq := NewDeque[int]()
			for _, query := range testCase.queries {
				switch query.queryType {
				case "EnQueueFront":
					dq.EnQueueFront(query.parameter)
				case "EnQueueRear":
					dq.EnQueueRear(query.parameter)
				case "DeQueueFront":
					result, err := dq.DeQueueFront()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "DeQueueRear":
					result, err := dq.DeQueueRear()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "Front":
					result, err := dq.Front()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "Rear":
					result, err := dq.Rear()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "IsEmpty":
					result := dq.Empty()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Length":
					result := dq.Length()
					if result != query.expectedResult {
						t.Errorf("Expected %v got %v", query.expectedResult, result)
					}
				}
			}
		})
	}

	// Run tests with strings
	for _, testCase := range testCasesString {
		t.Run(testCase.description, func(t *testing.T) {
			dq := NewDeque[string]()
			for _, query := range testCase.queries {
				switch query.queryType {
				case "EnQueueFront":
					dq.EnQueueFront(query.parameter)
				case "EnQueueRear":
					dq.EnQueueRear(query.parameter)
				case "DeQueueFront":
					result, err := dq.DeQueueFront()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "DeQueueRear":
					result, err := dq.DeQueueRear()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "Front":
					result, err := dq.Front()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "Rear":
					result, err := dq.Rear()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
				case "IsEmpty":
					result := dq.Empty()
					if result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Length":
					result := dq.Length()
					if result != query.expectedResult {
						t.Errorf("Expected %v got %v", query.expectedResult, result)
					}
				}
			}
		})
	}
}
