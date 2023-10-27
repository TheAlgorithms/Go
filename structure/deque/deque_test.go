//
// This file contains unit tests for the deque package.
// The tests cover the following scenarios:
// - Empty deque
// - Deque with one element
// - Deque with multiple elements
// The tests are parameterized with int and string types.
// Each test case is defined with a description and a list of queries to be executed on the deque.
// The expected results and errors are also defined for each query.
//

package deque_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/structure/deque"
)

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
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:     "Front",
					expectedError: deque.ErrEmptyDequeue,
				},
				{
					queryType:     "Rear",
					expectedError: deque.ErrEmptyDequeue,
				},
				{
					queryType:     "DeQueueFront",
					expectedError: deque.ErrEmptyDequeue,
				},
				{
					queryType:     "DeQueueRear",
					expectedError: deque.ErrEmptyDequeue,
				},
				{
					queryType:      "Length",
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test deque with one element",
			queries: []QueryStructure[int]{
				{
					queryType:     "EnQueueFront",
					parameter:     1,
					expectedError: nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test deque with multiple elements",
			queries: []QueryStructure[int]{
				{
					queryType:     "EnQueueFront",
					parameter:     1,
					expectedError: nil,
				},
				{
					queryType:     "EnQueueFront",
					parameter:     2,
					expectedError: nil,
				},
				{
					queryType:     "EnQueueRear",
					parameter:     3,
					expectedError: nil,
				},
				{
					queryType:     "EnQueueRear",
					parameter:     4,
					expectedError: nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					expectedResult: 2,
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					expectedResult: 2,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueRear",
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
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
					queryType:     "EnQueueFront",
					parameter:     "a",
					expectedError: nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 1,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					expectedResult: "a",
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: true,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 0,
					expectedError:  nil,
				},
			},
		},
		{
			description: "Test multiple elements deque",
			queries: []QueryStructure[string]{
				{
					queryType:     "EnQueueFront",
					parameter:     "a",
					expectedError: nil,
				},
				{
					queryType:     "EnQueueFront",
					parameter:     "b",
					expectedError: nil,
				},
				{
					queryType:     "EnQueueRear",
					parameter:     "c",
					expectedError: nil,
				},
				{
					queryType:     "EnQueueRear",
					parameter:     "d",
					expectedError: nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Front",
					expectedResult: "b",
					expectedError:  nil,
				},
				{
					queryType:      "Rear",
					expectedResult: "d",
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 4,
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueFront",
					expectedResult: "b",
					expectedError:  nil,
				},
				{
					queryType:      "DeQueueRear",
					expectedResult: "d",
					expectedError:  nil,
				},
				{
					queryType:      "IsEmpty",
					expectedResult: false,
					expectedError:  nil,
				},
				{
					queryType:      "Length",
					expectedResult: 2,
					expectedError:  nil,
				},
			},
		},
	}

	// Run tests with ints
	for _, testCase := range testCasesInt {
		t.Run(testCase.description, func(t *testing.T) {
			dq := deque.New[int]()
			for _, query := range testCase.queries {
				switch query.queryType {
				case "EnQueueFront":
					dq.EnqueueFront(query.parameter)
				case "EnQueueRear":
					dq.EnqueueRear(query.parameter)
				case "DeQueueFront":
					result, err := dq.DequeueFront()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "DeQueueRear":
					result, err := dq.DequeueRear()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Front":
					result, err := dq.Front()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v, %v", query.expectedResult, result, testCase.description)
					}
				case "Rear":
					result, err := dq.Rear()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "IsEmpty":
					result := dq.IsEmpty()
					if result != query.expectedResult {
						t.Errorf("Expected error: %v, got : %v", query.expectedResult, result)
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
			dq := deque.New[string]()
			for _, query := range testCase.queries {
				switch query.queryType {
				case "EnQueueFront":
					dq.EnqueueFront(query.parameter)
				case "EnQueueRear":
					dq.EnqueueRear(query.parameter)
				case "DeQueueFront":
					result, err := dq.DequeueFront()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "DeQueueRear":
					result, err := dq.DequeueRear()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Front":
					result, err := dq.Front()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v, %v", query.expectedResult, result, testCase.description)
					}
				case "Rear":
					result, err := dq.Rear()
					if err != query.expectedError {
						t.Errorf("Expected error: %v, got : %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "IsEmpty":
					result := dq.IsEmpty()
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
