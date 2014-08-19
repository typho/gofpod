package objects

import "io"

/*PdfValueType const (
    array
    boolean
    dictionary
    integer
    name
    null
    real
    stream
    string
)*/

type PdfType interface {
    Read(reader *io.ByteReader)
    Write(writer *io.ByteWriter)
}

