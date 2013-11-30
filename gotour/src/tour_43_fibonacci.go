package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    work := 0
    prev := 0
    sum := 1
    return func() int {
        work = prev
        prev = sum
        sum += work
        return sum
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 20; i++ {
        fmt.Println(f())
    }
}
