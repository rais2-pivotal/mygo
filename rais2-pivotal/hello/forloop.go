package main

import "fmt"

func main() {
   scores := [8]float64{10.30, 34.22, 100, 26.22, 88.4, 92.65, 32.11, 77.62}
   
   avg := 0.0
 
   for i := 0; i<8; i++ {
      avg += scores[i]
   }

   avg = avg / 8

   fmt.Printf("Average of array is %.2f\n", avg)

}


