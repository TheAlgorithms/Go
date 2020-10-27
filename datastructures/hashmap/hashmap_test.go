package hashmap_test

import (
	"TheAlgorithms/Go/datastructures/hashmap"
	"fmt"
	"testing"
)

func TestHashMap_Contains(t *testing.T) {

	mp := hashmap.New()

	mp.Put("test-1", 10)
	fmt.Println(mp.Get("test-1"))

	mp.Put("test-1", 20)
	mp.Put("test-2", 30)
	mp.Put(1, 40)

	fmt.Println(mp.Get("test-1"))
	fmt.Println(mp.Get("test-2"))
	fmt.Println(mp.Get(1))

	fmt.Println(mp.Contains(2))
	fmt.Println(mp.Contains(1))
	fmt.Println(mp.Contains("test-1"))

}
