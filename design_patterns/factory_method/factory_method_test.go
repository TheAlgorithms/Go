package factory_method

import (
	"reflect"
	"testing"
)

// test client code and how  to use the factory
func TestFactoryMethod(t *testing.T) {
	t.Run("Test Accounting Department", func(t *testing.T) {
		accounting, _ := getDepartment("accounting")
		if !reflect.DeepEqual(accounting.getName(), "Accounting Department") {
			t.Errorf("got: %v, want: %v", accounting.getName(), "Accounting Department")
		}
		if !reflect.DeepEqual(accounting.getEmployees(), 1) {
			t.Errorf("got: %v, want: %v", accounting.getName(), 1)
		}
	})


	t.Run("Test Finance Department", func(t *testing.T) {
		finance, _ := getDepartment("finance")
		if !reflect.DeepEqual(finance.getName(), "Finance Department") {
			t.Errorf("got: %v, want: %v", finance.getName(), "Finance Department")
		}
		if !reflect.DeepEqual(finance.getEmployees(), 2) {
			t.Errorf("got: %v, want: %v", finance.getEmployees(), 2)
		}
	})

	t.Run("Test not existing department", func(t *testing.T) {
		random, _ := getDepartment("test")

		if !reflect.DeepEqual(random, nil) {
			t.Errorf("should return empty nil for non existing department")
		}

	})
}




