// Tests for the deque package.

package deque

import "testing"

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
					expectedError: ErrEmptyDeQueue,
				},
				{
					queryType:     "Rear",
					expectedError: ErrEmptyDeQueue,
				},
				{
					queryType:     "DeQueueFront",
					expectedError: ErrEmptyDeQueue,
				},
				{
					queryType:     "DeQueueRear",
					expectedError: ErrEmptyDeQueue,
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
			dq := NewDeque[int]()
			for _, query := range testCase.queries {
				switch query.queryType {
				case "EnQueueFront":
					dq.EnQueueFront(query.parameter)
				case "EnQueueRear":
					dq.EnQueueRear(query.parameter)
				case "DeQueueFront":
					result, err := dq.DeQueueFront()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "DeQueueRear":
					result, err := dq.DeQueueRear()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Front":
					result, err := dq.Front()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v, %v", query.expectedResult, result, testCase.description)
					}
				case "Rear":
					result, err := dq.Rear()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
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
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "DeQueueRear":
					result, err := dq.DeQueueRear()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
					}
				case "Front":
					result, err := dq.Front()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v, %v", query.expectedResult, result, testCase.description)
					}
				case "Rear":
					result, err := dq.Rear()
					if err != query.expectedError {
						t.Errorf("Expected %v, got %v", query.expectedError, err)
					}
					if err == nil && result != query.expectedResult {
						t.Errorf("Expected %v, got %v", query.expectedResult, result)
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
