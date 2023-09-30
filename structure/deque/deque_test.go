// Tests for the deque package.

package deque

import (
	"testing"
)

func TestDeque(t *testing.T) {
	d := newDeque()

	// Test empty deque
	if !d.isEmpty() {
		t.Errorf("Expected deque to be empty, but it's not")
	}

	// Test enQueue front
	d.enQueueFront(1)
	if d.isEmpty() {
		t.Errorf("Expected deque to not be empty, but it is")
	}
	val, err := d.getFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected front value to be 1, but got %v", val)
	}
	val, err = d.getRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected rear value to be 1, but got %v", val)
	}

	// Test enQueue rear
	d.enQueueRear(2)
	if d.isEmpty() {
		t.Errorf("Expected deque to not be empty, but it is")
	}
	val, err = d.getFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected front value to be 1, but got %v", val)
	}
	val, err = d.getRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 2 {
		t.Errorf("Expected rear value to be 1, but got %v", val)
	}

	// Test deQueue front
	val, err = d.deQueueFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected dequeued front value to be 1, but got %v", val)
	}
	if d.isEmpty() {
		t.Errorf("Expected deque to not be empty, but it is")
	}
	val, err = d.deQueueFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 2 {
		t.Errorf("Expected front value to be 2, but got %v", val)
	}

	// Test deQueue back
	d.enQueueFront(1)
	d.enQueueFront(2)

	val, err = d.deQueueRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected dequeued front value to be 1, but got %v", val)
	}
	if d.isEmpty() {
		t.Errorf("Expected deque to not be empty, but it is")
	}
	val, err = d.deQueueRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 2 {
		t.Errorf("Expected front value to be 2, but got %v", val)
	}

	// Test getFront
	_, err = d.getFront()
	if err == nil {
		t.Errorf("Expected empty queue error, but got none")
	}
	d.enQueueFront(1)
	val, err = d.getFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected front value to be 1, but got %v", val)
	}
	_, err = d.deQueueFront()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	// Test getRear
	_, err = d.getRear()
	if err == nil {
		t.Errorf("Expected empty queue error, but got none")
	}
	d.enQueueRear(1)
	val, err = d.getRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected rear value to be 1, but got %v", val)
	}
	_, err = d.deQueueRear()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	// Test empty deque
	if !d.isEmpty() {
		t.Errorf("Expected deque to be empty, but it's not")
	}

	// Test getLength
	if d.getLength() != 0 {
		t.Errorf("Expected deque length to be 0, but got %v", d.getLength())
	}
	d.enQueueFront(1)
	if d.getLength() != 1 {
		t.Errorf("Expected deque length to be 1, but got %v", d.getLength())
	}
	d.enQueueFront(2)
	if d.getLength() != 2 {
		t.Errorf("Expected deque length to be 2, but got %v", d.getLength())
	}

}
