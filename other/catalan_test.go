// reference for testing https://www.calhoun.io/how-to-test-with-go/
package testing101

import (

  "testing"
)

func TestCatalan(t *testing.T){

  x:=5
  expected:=42
  actual:= Catalan(x)

  if actual != expected {
    t.Errorf("Expected the sum of %d to be %d but instead got %d", x, expected, actual)
  }
}
