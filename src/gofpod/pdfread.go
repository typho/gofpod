package main

import (
    "bufio"
    "fmt"
    "os"
    "gofpod/objects"
)

func main() {
    fmt.Printf("Starting\n")

    // Reading PDF into fi
    fi, err := os.Open("example_pdfs/hello_world.pdf")
    if err != nil { panic(err) }
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // Create a reader for fi
    r := bufio.NewReader(fi)

    objects.ReadPdfFileHeader(r)
}

