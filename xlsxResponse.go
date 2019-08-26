package httpserve

import (
	"encoding/xml"
	"io"

	"github.com/tealeg/xlsx"
)

// NewXLSXResponse will return a new text response
func NewXLSXResponse(code int, value *xlsx.File) *XLSXResponse {
	var x XLSXResponse
	x.code = code
	x.val = value
	return &x
}

// XLSXResponse is a basic text response
type XLSXResponse struct {
	code int
	val  *xlsx.File
}

// ContentType returns the content type
func (x *XLSXResponse) ContentType() (contentType string) {
	return "application/xlsx"
}

// StatusCode returns the status code
func (x *XLSXResponse) StatusCode() (code int) {
	return x.code
}

// WriteTo will write to a given io.Writer
func (x *XLSXResponse) WriteTo(w io.Writer) (n int64, err error) {
	var mp map[string]string
	// Initialize a new XLSX value
	if mp, err = x.val.MarshallParts(); err != nil {
		return
	}
	// Initialize a new XLSX encoder
	enc := xml.NewEncoder(w)
	// Encode the responder
	err = enc.Encode(mp)
	return
}
