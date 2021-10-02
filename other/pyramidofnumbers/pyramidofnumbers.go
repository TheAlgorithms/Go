// Program in Golang to print Pyramid of Numbers
/* author Sreejita Roy (https://github.com/pseudonerd16)*/

package main

import "fmt"

func main() {
	var rows,k,temp,temp1 int
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		
		for j := 1; j <= rows-i; j++ {
			fmt.Print("  ")
			temp++
		}
		for{
			if( temp <= rows-1){
				fmt.Printf(" %d",i+k)
				temp++
			}else{
				temp1++
				fmt.Printf(" %d",(i+k-2*temp1))
			}
			k++

			if(k == 2*i-1){				
				break
			}

		}
		temp = 0
		temp1 = 0
		k = 0
		fmt.Println("")
	}

}