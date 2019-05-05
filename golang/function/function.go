package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(eval(3, 5, "*"))
	fmt.Println(eval(3, 5, "x"))
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
    fmt.Println(sum(1, 2, 3))
    fmt.Println(calcTriangle(4, 3))
}

func calcTriangle(a, b int) int {
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))
    return c
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s "+"with args: %d, %d\n", opName, a, b)
	return op(a, b)
}

func sum(a ...int) int {
    var b = 0
    for _, i := range a {
        b += i
    }
    return b
}
