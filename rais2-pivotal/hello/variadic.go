package main

import "fmt"

func main() {
   a, b, c, d, e := 1, 2, 3, 4, 5
   sum := add(a,b,c,d)
   fmt.Printf("Value of sum is %v\n", sum)
   a = e
}

func add(ints ...int) int {
   sum := 0
   for _, v := range ints {
      fmt.Printf("Value of v is %v\n", v)
      sum += v
   }
   return sum
}

