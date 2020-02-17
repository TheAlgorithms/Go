package ArrayList

import "fmt"
import "strconv"
import "strings"
import "math/rand"
import (
	"time"
)

type List struct {
	size int
	pointer int
	array []int
}

func NewList() *List {
	var array = make([]int, 8)
	l := List{pointer: 0, size: 8, array: array}

	return &l
} 

func (l *List) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(l.Length(), func(i, j int) {
		l.Swap(i, j)
	})
} 

func (l *List) grow() {
	var newSize = l.size * 2

	var array = make([]int, newSize)
	for i := 0; i < l.size; i++ {
		array[i] = l.array[i]
	}
	l.array = array
	l.size = newSize
	
}

func (l *List) GetIndex(index int) int {
	return l.array[index]
}

func (l *List) Append(item int) {

	if l.pointer < l.size {
		l.array[l.pointer] = item
		l.pointer += 1
	}else {
		l.grow()
		l.Append(item)
	}

}

func (l *List) Print() {
	var out = "["
	for i := 0; i<l.pointer; i++{
		out += strconv.Itoa(l.array[i])
		out += ", "
	
	}
	out = strings.TrimSuffix(out, ", ")
	out += "]"
	fmt.Println(out)
}

func (l *List) Swap(index1 int, index2 int) {

	l.array[index1], l.array[index2] = l.array[index2], l.array[index1]

}

func (l *List) Length() int {
	return l.pointer
}

func (l *List) Pop() int {
	l.pointer -= 1
	return l.array[l.pointer]
}

// func main(){
// 	var l = List()
// 	l.append(2)
// 	l.append(4)
// 	l.print()
// 	for i:=5; i < 20; i++ {
// 		l.append(i)
// 	}
// 	l.print()
// 	l.swap(0, 1)
// 	l.print()
// }