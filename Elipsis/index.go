package main

import "fmt"

func perkalianDua(angka ...int) []int {
   hasil := []int{}
   for _, num := range angka {
      cal := num * 2
      hasil = append(hasil, cal)
   }
   return hasil
}

func main() {
   fmt.Println(perkalianDua(2))         // Output: [4]
   fmt.Println(perkalianDua(1, 2, 3, 4, 5)) // Output: [2 4 6 8 10]
}