package main

import "fmt"

func main() {
    source := []int{1, 2, 3, 4, 5}
    destination := make([]int, 10)

    copy(destination, source)

    fmt.Println(destination)
}
