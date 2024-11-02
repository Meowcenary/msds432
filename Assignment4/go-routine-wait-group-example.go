package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // example1()
    example2()
}

func example1() {
    var wg sync.WaitGroup
    salutation := "hello"
    wg.Add(1)
    go func() {
        defer wg.Done()
        salutation = "welcome"
    }()
    fmt.Println(salutation)
    wg.Wait()
    fmt.Println(salutation)
}

func example2() {
    var wg sync.WaitGroup
    for _, salutation := range []string{"hello", "greetings", "good day"} {
        wg.Add(1)
        go func() {
            defer wg.Done()
            time.Sleep(8 * time.Second)
            fmt.Println(salutation)
        }()
    }
    wg.Wait()
}
