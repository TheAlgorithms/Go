package prototype

import (
	"fmt"
	"testing"
)

// test client code and how  to use the factory
func TestProtoType(t *testing.T) {
	t.Run("Test File And Folder Hierarchy", func(t *testing.T) {
		file1 := &file{name: "File1"}
		file2 := &file{name: "File2"}
		file3 := &file{name: "File3"}

		folder1 := &folder{
			children: []nodeInterface{file1},
			name:      "Folder1",
		}

		folder2 := &folder{
			children: []nodeInterface{folder1, file2, file3},
			name:      "Folder2",
		}
		fmt.Println("\nPrinting hierarchy for Folder2")
		folder2.print("  ")

		cloneFolder := folder2.clone()
		fmt.Println("\nPrinting hierarchy for clone Folder")
		cloneFolder.print("  ")
	})
}





