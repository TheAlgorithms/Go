package cache

import "testing"

// TestLFU1 tests the basic functionality of the LFU cache
// input : ["NewLFU(2)", "put(one,1)", "put(two,2)", "get("one")", "put("three,3")","get("two")", "get("three")", "put("four",4)", "get("one")", "get("three")", "get("four")""]
// output : [nil, nil, nil, 1, nil, nil, 3, nil, nil, 3, 4]
func TestLFU1(t *testing.T) {
	cache := NewLFU(2)

	cache.Put("one", 1)
	cache.Put("two", 2)

	val, err := cache.Get("one")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 1 {
		t.Error("Expected 1")
	}

	cache.Put("three", 3)

	val, err = cache.Get("two")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}

	val, err = cache.Get("three")

	if err != nil {
		t.Error("Expected no error")
	}
	if val != 3 {
		t.Error("Expected 3")
	}

	cache.Put("four", 4)

	val, err = cache.Get("one")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}

	val, err = cache.Get("three")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 3 {
		t.Error("Expected 3")
	}

	val, err = cache.Get("four")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 4 {
		t.Error("Expected 4")
	}

}

// TestLFU2 tests the corner case when the capacity is 0
func TestLFU2(t *testing.T) {
	cache := NewLFU(0)
	cache.Put("0", 0)
	val, err := cache.Get("0")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}
}

// TestLfU3 tests the case
// input : ["NewLFU(3)", "put(one,1)", "put(two,2)", "put(three,3)", "put(four,4)", "get(four)", "get(three)", "get(two)", "get(one)", "put(five,5)", "get(one)", "get(two)", "get(three)", "get(four)", "get(five)"]
// output : [nil, nil, nil, nil, nil, 4, 3, 2, nil, nil, nil, 2, 3, nil, 5]
func TestLFU3(t *testing.T) {
	cache := NewLFU(3)
	cache.Put("one", 1)
	cache.Put("two", 2)
	cache.Put("three", 3)
	cache.Put("four", 4)

	val, err := cache.Get("four")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 4 {
		t.Error("Expected 4")
	}
	val, err = cache.Get("three")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 3 {
		t.Error("Expected 3")
	}
	val, err = cache.Get("two")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 2 {
		t.Error("Expected 2")
	}

	val, err = cache.Get("one")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}

	cache.Put("five", 5)

	val, err = cache.Get("one")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}

	val, err = cache.Get("two")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 2 {
		t.Error("Expected 2")
	}

	val, err = cache.Get("three")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 3 {
		t.Error("Expected 3")
	}

	val, err = cache.Get("four")
	if err == nil {
		t.Error("Expected error")
	}
	if val != nil {
		t.Error("Expected nil")
	}

	val, err = cache.Get("five")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != 5 {
		t.Error("Expected 5")
	}

}

// TestLFU4 tests the the case item is of string type
func TestLFU4(t *testing.T) {
	cache := NewLFU(2)
	cache.Put("two", "one")
	cache.Put("two", "two")
	val, err := cache.Get("two")
	if err != nil {
		t.Error("Expected no error")
	}
	if val != "two" {
		t.Error("Expected two")
	}
}
