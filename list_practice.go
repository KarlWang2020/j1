
package main

import (
	"container/list"
	"fmt"
)


func main() {
	wel:=list.New()
	wel.PushBack(12)
	element1:=wel.PushFront(1212)
	wel.PushBack(344343)
	wel.InsertAfter(155552,element1)
	wel.PushBack(121)
	for i:=wel.Front(); i!=nil ; i=i.Next() {
           fmt.Println(i.Value)
           if i.Value==12{
           	wel.Remove(i)
		   }

	}




  fmt.Println()
	for i:=wel.Front(); i!=nil ; i=i.Next() {
		fmt.Println(i.Value)
		if i.Value==12{
			n:=i
			i=i.Next()
			wel.Remove(n)
			continue
		}

	}
}

