package main

import (
    "fmt"
    "os"
    "strconv"
)


func fib(n int) chan int{
    c := make(chan int)
    go func() {
        a, b := 0, 1
        for i := 0; i < n; i++ {
            a, b =  b, a+b
            c <- a
        }
        close(c)
    }()
    return c
}


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Integer argument required.")
        os.Exit(-1)
    }
    number, err := strconv.Atoi(os.Args[1])
    if err != nil || number <= 0 {
        fmt.Println("Integer argument greater than 0 required.")
        os.Exit(-1)
    }
    for x := range fib(number) {
        fmt.Println(x)
    }
}
