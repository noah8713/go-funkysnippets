// program to find square root of a number
// Solves: go tutorial https://tour.golang.org/flowcontrol/8
// run as: go run sqrt.go
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    t, z := 0., 1.
    for {
        z, t = z - (z * z-x ) / (2 * z), z
        fmt.Println("z is:", z)
        if math.Abs(t-z) < 1e-8 {
            break
        }
    }
    return z
}

func main() {
    i := 3.
    sqrt_i := Sqrt(i)
    fmt.Println(sqrt_i)
    fmt.Println(sqrt_i == math.Sqrt(i))
}
