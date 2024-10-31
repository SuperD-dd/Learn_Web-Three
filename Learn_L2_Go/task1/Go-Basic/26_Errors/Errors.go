// package main

// import (
//     "errors"
//     "fmt"
// )

// func f(arg int) (int, error) {
//     if arg == 42 {

//         return -1, errors.New("can't work with 42")
//     }

//     return arg + 3, nil
// }

// var ErrOutOfTea = fmt.Errorf("no more tea available")
// var ErrPower = fmt.Errorf("can't boil water")

// func makeTea(arg int) error {
//     if arg == 2 {
//         return ErrOutOfTea
//     } else if arg == 4 {

//         return fmt.Errorf("making tea: %w", ErrPower)
//     }
//     return nil
// }

// func main() {
//     for _, i := range []int{7, 42} {

//         if r, e := f(i); e != nil {
//             fmt.Println("f failed:", e)
//         } else {
//             fmt.Println("f worked:", r)
//         }
//     }

//     for i := range 5 {
//         if err := makeTea(i); err != nil {

//             if errors.Is(err, ErrOutOfTea) {
//                 fmt.Println("We should buy new tea!")
//             } else if errors.Is(err, ErrPower) {
//                 fmt.Println("Now it is dark.")
//             } else {
//                 fmt.Printf("unknown error: %s\n", err)
//             }
//             continue
//         }

//         fmt.Println("Tea is ready!")
//     }
// }

package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {

    c.mu.Lock()
    c.counters[name]++
    defer c.mu.Unlock()
}

func main() {
    c := Container{

        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()
    fmt.Println(c.counters)
}