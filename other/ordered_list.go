package main
import ("fmt")

// function to remove an element in an ordered list in sequential allocation
// https://github.com/lucastrogo
var pos int
var key int
var listNum int

func remove(listNum[]int) int {
  pos = 0		// pos of the element
  searcher := len(listNum) - 1   // variable to search the element
  key = 2		  // where is located
  for(pos <= searcher){   // looping for search the element and remove
    if key == pos{
      if key == searcher{
        listNum[key] = listNum[key - 1]
        pos ++}else{
        listNum[key] = listNum[key + 1]
        pos ++
      }
      }else{pos++}
  }
  return listNum[key]
}

func main(){
	listNum := []int{1, 2, 3, 4, 5}
	fmt.Println(listNum)
}