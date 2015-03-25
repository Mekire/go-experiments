package main

import (
    "fmt"
    "os"
    "strconv"
)


func fib() chan int{
    c := make(chan int)
    go func() {
        a, b := 0, 1
        for {
            a, b =  b, a+b
            c <- a
        }
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
    generator := fib()
    for i := 0 ; i < number; i++{
        fmt.Println(<-generator)
    }
    close(generator)
}
