package Heap

import (
	"fmt"
	array "../arrayList"
)


type Heap struct {
	list *array.List 
}

func NewHeap() *Heap {
	var heap = Heap{list: array.NewList()}
	return &heap
}

func (h *Heap) Print(){
	h.list.Print()
}

func (h *Heap) Add(value int) {
	h.list.Append(value)
	h.bubbleUp(h.Length() - 1)
}

func (h *Heap) bubbleUp(index int) {
	var parent = (index - 1) / 2

	if h.list.GetIndex(index) < h.list.GetIndex(parent) {
		h.list.Swap(index, parent)
		h.bubbleUp(parent)
	}

}

func (h *Heap) RemoveMin() int {
	var returnValue = h.min()

	h.list.Swap(0, h.Length()-1)

	h.list.Pop()

	h.bubbleDown(0)

	return returnValue
}

func (h *Heap) bubbleDown(index int) {
	var left = (2*index) + 1
	var right = (2*index) + 2
	if left < h.Length() || right < h.Length() {
		var target = left
		if right < h.Length() && h.list.GetIndex(right) < h.list.GetIndex(left) {
			target = right
		}
		if h.list.GetIndex(index) > h.list.GetIndex(target) {
			h.list.Swap(index, target)
			h.bubbleDown(target)
		}
	}
}

func (h *Heap) Length() int {
	return h.list.Length()
}

func (h *Heap) min() int {
	return h.list.GetIndex(0)
}

func main(){
	fmt.Println("Hello world")
}