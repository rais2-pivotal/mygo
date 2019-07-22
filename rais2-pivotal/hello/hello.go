package main

import "fmt"

var (
    a = 10 //inferred type (without specifiying type like int
    b = 20 
    
    c int = 10 //declared type
    d int = 20 

    e int = 10; f int = 20
    
    g,h = 10,20
    
    )

func main() {

        i,j := 10, 20
	fmt.Printf("hello, world\n")
        fmt.Printf("%v + %v = %v\n",a ,b , a + b)
        fmt.Printf("%v + %v = %v\n",c ,d , c + d)
        fmt.Printf("%v + %v = %v\n",e ,f , e + f)
        fmt.Printf("%v + %v = %v\n",g ,h , g + h)
        fmt.Printf("%v + %v = %v\n",i ,j , i + j)
}
