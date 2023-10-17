package graph

import (
	"testing"
)

func TestUnionFind(t *testing.T){
    // create a union find data structure with 10 elements
	u:= NewUnionFind(10)
	// union operations 
	u = u.Union(0,1)
	u = u.Union(2,3)
	u = u.Union(4,5)
	u = u.Union(6,7)
	

// test finding the parent of specific elements
t.Run("Test Find", func(t *testing.T){
	if(u.Find(0)!= u.Find(1)!= u.Find(2)!= u.Find(3)!= u.Find(4)!= u.Find(5)!= u.Find(6)!= u.Find(7)){
		t.Error("Union operation not functioning correctly")
	}
})

// more union operations
   u = u.Union(1,5)
   u = u.Union(3,7)

// test finding the parent of specific elements after more union operations
t.Run("Test Find after union", func(t *testing.T){
	if(u.Find(0)!= u.Find(5)!= u.Find(2)!= u.Find(7)){
		t.Error("Union operation not functioning correctly")
	}
})
}