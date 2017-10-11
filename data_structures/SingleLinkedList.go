package main

import "fmt"

type Item struct { 
    nextNode *Item
    val int
}

// initilizes the list
func initList() *Item {
    pHead := &Item{nil, 0}  
    return pHead
}

// prints the list
func printList(pList *Item) {
    pCurr := pList
    for {
        fmt.Printf("%d ", pCurr.val)
        
        if pCurr.nextNode != nil {
            pCurr = pCurr.nextNode
        } else {
            break
        }
    }
    fmt.Println("")
}

// inserts element at the beginning of the list
func prepend(pList *Item, val int) *Item {
  pItem := &Item{nil, val}
  pItem.nextNode = pList
  pList = pItem
  return pList
}

// appends element at the end of the list
func append(pList *Item, val int) *Item {
  pCurr := pList
  for {
      if pCurr.nextNode != nil {
          pCurr = pCurr.nextNode
      } else {
          pItem := &Item{nil, val}
          pCurr.nextNode = pItem
          break
      }
  }
  return pCurr
}

// insert a new node after the given prevNode
func insertAfter(prevNode *Item, val int) {
  if prevNode == nil {
    fmt.Println("previous node cannot be NIL")
  }
  pItem := &Item{nil, val}
  pItem.nextNode = prevNode.nextNode
  prevNode.nextNode = pItem
}

// removes an element from the beginning of the list
func removeHead(pList *Item) *Item {
  if pList != nil {
    pList = pList.nextNode
  }
  return pList
}

// removes an element from the end of the list
func removeTail(pList *Item) *Item {
  pCurr := pList
  for {
      if (pCurr.nextNode).nextNode != nil {
          pList = pCurr
          pCurr = pCurr.nextNode
      } else {
          pCurr.nextNode = nil
          break
      }
  }
  return pList
}

// reverses the list
func reverseList(pList *Item) *Item {
  pCurr := pList
  var pTop *Item = nil
  for {
      if pCurr == nil {
          break
      }
      pTemp := pCurr.nextNode
      pCurr.nextNode = pTop
      pTop = pCurr
      pCurr = pTemp        
  }
  
  return pTop
}

func main() {
  pList := initList()
  pList = append(pList, 1)
  pList = append(pList, 3)
  printList(pList) // prints 1 3
  pList = prepend(pList, 0)
  printList(pList) // prints 0 1 3
  insertAfter(pList.nextNode, 2)
  printList(pList) // prints 0 1 2 3
  pList = removeHead(pList)
  printList(pList) // prints 1 2 3
  pList = removeTail(pList)
  printList(pList) // prints 1 2
  pList = reverseList(pList)
  printList(pList) // prints 2 1
}
