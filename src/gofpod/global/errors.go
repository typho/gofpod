package main

import (
    "fmt"
    "strconv"
)

// A SyntaxError is returned when the PDF is syntactically invalid
type SyntaxError struct {
    what string    // message indicating the context
    where int      // byte index
}

func (e SyntaxError) Error() string {
    return e.what
}

func ret() (int, error) {
    if false {
        return 1, nil
    }
    return 0, SyntaxError{"Hell", 150}
}

func main() {
    if res, err := ret(); err != nil {
        fmt.Println("A SyntaxError was thrown: " + err.Error());
    } else {
        fmt.Println("Hello World: " + strconv.Itoa(res))
    }
}
