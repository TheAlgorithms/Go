// Code for Ugly Numbers (Numbers with prime factors 2,3 and 5)
// Ref: https://www.geeksforgeeks.org/ugly-numbers/
// https://leetcode.com/problems/ugly-number/

package main
import "fmt"

func maxDivide(a, b int)int{
    for (a%b==0){
        
        a=a/b
    }
    return a
}

func isUgly(no int)bool{

    
    no= maxDivide(no, 2)
    no= maxDivide(no, 3)
    no= maxDivide(no, 5)
    
    if (no==1){
        
        return true
    }else{
        return false
    }
}


func getNthUglyNo(n int)int{
    
    var i int = 1
    var count int = 1

    
    for (n>count){
        
        i++
        
        if (isUgly(i)){
            count++
        }
    }  
    return i
}


func main(){
    
    no := 150
    result := getNthUglyNo(no)
    converted_result := uint(result)
    fmt.Println("150th ugly number is", converted_result)
    
}