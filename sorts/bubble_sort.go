//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func bubbleSort(array []int)[]int{
	for a:=0; a<len(array)-1;a++{
		for i:=a;i<len(array)-1;i++{
			j:=i+1
			array =swap(array, i,j)
		}
	}
	return array

}

func swap(array []int, i int, j int)[]int{
	if array[i]>array[j]{
		array[i],array[j] =array[j],array[i]
	}
	return array

}