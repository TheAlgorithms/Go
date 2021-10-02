package main
import "fmt"
/**
 * Program to implement Kadane’s Algorithm to calculate maximum contiguous subarray sum of an array
 * Time Complexity: O(n)
 *
 * @author Ritik Agarwal
*/
 /**
   * This method implements Max operation
   *
   * @param a,b Two integers
   * @return The maximum among the two integers
   */
func max(a int, b int) int{
    if a>b{
        return a
    }
    return b
}
 /**
   * This method implements Kadane's Algorithm
   *
   * @param nums The input array
   * @return The maximum contiguous subarray sum of the array
   */
func maxSubArray(nums []int) int {
    if(len(nums)==0){
        return 0
    }
    current_sum:=nums[0]
    max_sum:=nums[0]
    for i:=1;i<len(nums);i++{
        current_sum=max(current_sum+nums[i],nums[i])
        if(current_sum>max_sum){
            max_sum=current_sum
        }
        if(current_sum<0){
            current_sum=0
        }
    }
    return max_sum
}
  /**
   * Main method
   *
   */
func main(){
   var size int
   fmt.Scanln(&size)
   var arr = make([]int, size)
   for i:=0; i<size; i++ {
      fmt.Scanf("%d", &arr[i])
   }
   fmt.Println(maxSubArray(arr))
}
