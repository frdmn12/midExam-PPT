package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var array1 [3]string
    var array2 [3]string

    // Input array 1
    fmt.Println("Masukkan isi array 1, dipisahkan oleh spasi:")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        values := strings.Split(input, " ")
        for i := 0; i < len(values); i++ {
            array1[i] = values[i]
        }
    }

    // Input array 2
    fmt.Println("Masukkan isi array 2, dipisahkan oleh spasi:")
    if scanner.Scan() {
        input := scanner.Text()
        values := strings.Split(input, " ")
        for i := 0; i < len(values); i++ {
            array2[i] = values[i]
        }
    }

    // Compare arrays
    for i := 0; i < len(array1); i++ {
        if array1[i] != array2[i] {
            fmt.Printf("index ke %d berbeda\n", i)
        }
    }
}
