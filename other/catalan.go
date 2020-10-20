package testing101

func Catalan(x int) int{

  var  i int

  if x <=1{

    return 1
  }

  result:=0
  
  for i=0; i<x; i++ {

    result += Catalan(i)*Catalan(x-i-1)

  }
  return result
}
