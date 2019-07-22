package main

import "fmt"

func main() {
   fmt.Printf("5! = %v\n", factor(5))

}

func factor(n int) int {
   if n == 0 {
      return 1
   }

   return n * factor(n-1)
}
