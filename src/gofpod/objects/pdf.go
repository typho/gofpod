package objects

import (
    "io"
    "fmt"
)

func ReadPdfFileHeader(reader io.ByteReader) {
    var version uint8 = 8
    expected := []byte("%PDF-1.")

    for i := 0; i < 8; i++ {
        var byte, err = reader.ReadByte()
        if err != nil { panic(err) }
        if i < len(expected) {
            if byte != expected[i] {
                panic(err)
            }
        }
        if i == 7 {
            if 48 <= byte && byte <= 77 {
                version = uint8(byte - 48)
            }
        }
    }

    
    fmt.Println("PDF of version", version, "found")
}
