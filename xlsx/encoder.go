package xlsx

// Encoder is a basic XLSX encoder interface
type Encoder interface {
	MarshalXLSX(*File) error
}
