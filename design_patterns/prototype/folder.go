package prototype

import "fmt"

type folder struct {
	children []nodeInterface
	name      string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() nodeInterface {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []nodeInterface
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

