package main

import "fmt"

func swap(a,b int) (int,int) {
   tmp := a
   a = b 
   b = tmp
   return a,b
}

func swapsimple(a,b int) (int,int) {
   a,b = b,a
   return a,b
}

func swapnext(a,b int) (x,y int) {
   x, y = b, a
   return
}

func main() {
   v1,v2 := 3,5
   fmt.Printf("Value of v1 and v2 before swap v1 = %v and v2 = %v\n", v1, v2)
   //v1,v2 = swap(v1,v2)
   //v1,v2 = swapsimple(v1,v2)
   v1, v2 = swapnext(v1,v2)
   fmt.Printf("Value of v1 and v2 after swap v1 = %v and v2 = %v\n", v1, v2)
}
