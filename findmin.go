package main

import("fmt")

func findmin(in []int) int{
	smallest := in[0]
    for i := 0; i < len(in); i ++ {
        var x = in[i]

        if x <= smallest {
            smallest = x
        }

    }

    return smallest
}

func main() {
	slice := []int{1,2,3, 0, 5}
	fmt.Println(findmin(slice))
}