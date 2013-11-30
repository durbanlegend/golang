package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(f float64) (float64, error) {
    
    if f < 0 {
        return f, ErrNegativeSqrt(f)
    }

    
    z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - f) / (2 * z)
	}
	return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
