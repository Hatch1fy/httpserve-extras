package xlsx

import (
	"io"

	xlsx "github.com/Hatch1fy/tealeg-xlsx"
)

// NewResponse will return a new XLSX response
func NewResponse(code int, enc Encoder) *Response {
	var r Response
	r.code = code
	r.enc = enc
	return &r
}

// Response is an XLSX httpserve response response
type Response struct {
	code int
	enc  Encoder
}

// ContentType returns the content type
func (x *Response) ContentType() (contentType string) {
	return "application/xlsx"
}

// StatusCode returns the status code
func (x *Response) StatusCode() (code int) {
	return x.code
}

// WriteTo will write to a given io.Writer
func (x *Response) WriteTo(w io.Writer) (n int64, err error) {
	file := xlsx.NewFile()
	if err = x.enc.MarshalXLSX(file); err != nil {
		return
	}

	err = file.Write(w)
	return
}
