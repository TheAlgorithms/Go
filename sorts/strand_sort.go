package main
 
import "fmt"
 
type link struct {
    int
    next *link
}
 
func linkInts(s []int) *link {
    if len(s) == 0 {
        return nil
    }
    return &link{s[0], linkInts(s[1:])}
}
 
func (l *link) String() string {
    if l == nil {
        return "nil"
    }
    r := fmt.Sprintf("[%d", l.int)
    for l = l.next; l != nil; l = l.next {
        r = fmt.Sprintf("%s %d", r, l.int)
    }
    return r + "]"
}
 
func main() {
    a := linkInts([]int{170, 45, 75, -90, -802, 24, 2, 66})
    fmt.Println("before:", a)
    b := strandSort(a)
    fmt.Println("after: ", b)
}
 
func strandSort(a *link) (result *link) {
    for a != nil {
    
        // build sublist
        sublist := a
        a = a.next
        sTail := sublist
        for p, pPrev := a, a; p != nil; p = p.next {
            if p.int > sTail.int {
            
                // append to sublist
                sTail.next = p
                sTail = p
                
                // remove from a
                if p == a {
                    a = p.next
                } else {
                    pPrev.next = p.next
                }
            } else {
                pPrev = p
            }
        }
        sTail.next = nil 
        
        // terminate sublist
        if result == nil {
            result = sublist
            continue
        }
       
       // merge
        var m, rr *link
        if sublist.int < result.int {
            m = sublist
            sublist = m.next
            rr = result
        } else {
            m = result
            rr = m.next
        }
        result = m
        for {
            if sublist == nil {
                m.next = rr
                break
            }
            if rr == nil {
                m.next = sublist
                break
            }
            if sublist.int < rr.int {
                m.next = sublist
                m = sublist
                sublist = m.next
            } else {
                m.next = rr
                m = rr
                rr = m.next
            }
        }
    }
    return
}
