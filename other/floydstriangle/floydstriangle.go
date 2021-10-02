// Golang Program to Print Floyd's Triangle.
/* author Sreejita Roy (https://github.com/pseudonerd16)*/

package main

import "fmt"

func main(){
	var rows int
	var temp int = 1
	fmt.Print("Enter number of rows : ")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {	
		
		for k := 1; k <= i; k++ {

			fmt.Printf(" %d",temp)				
			temp++
		}
		fmt.Println("")		
	}

}