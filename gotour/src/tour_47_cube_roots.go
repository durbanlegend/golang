package main

import (
    "fmt"
    // "math"
    "math/cmplx"
    )

func Cbrt(x complex128) complex128 {
    z := x
    for i := 0; i < 10; i++ {
	    prev := z
    	z = z - (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
        delta := z - prev
        fmt.Printf("%v %v\n", z, delta)
        // if math.Abs((float64)delta) < 0.00001 {
        //     return z
        // }
        fmt.Println(delta)
    }
    return z
}

func main() {
    fmt.Println(Cbrt(27), cmplx.Pow(27, 1.0/3), cmplx.Pow(Cbrt(27), 3))
}