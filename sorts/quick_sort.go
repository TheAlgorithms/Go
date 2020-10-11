//Package sorts a package for demonstrating sorting algorithms in Go
package sorts


func quickSort(arr []int , first , last int) []int {
	i:= HoarePartition(arr,first,last)
	quickSort(arr, first, i)
	quickSort(arr, i+1, last)

}
func HoarePartition(arr []int , firsts, last int) int{
	pivot := arr[firsts]
	i := firsts - 1
	j := last +1

	for true{
		i++
		for arr[i] < pivot {
			i++
		}
		j--
		for arr[j]> pivot{
			j--
		}

		if i>=j{
			return j
		}
		arr[i],arr[j] = arr[j],arr[i]
	}
	return -1
}