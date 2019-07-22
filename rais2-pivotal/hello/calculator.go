package main

import "fmt"

func add(v1,v2 int) int {
   return v1 + v1
}

func main() {
   a, b := 3, 5
   fmt.Printf("Addition Output: ");
   fmt.Printf("%v\n", add(a,b));
}
