package main

import (
    "fmt"
    "math"
)

func main(){
    var a, b = 3, 4
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))
    fmt.Println(c)
}
