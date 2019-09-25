package binutil

// Reader is binary read utility.
// Note this interface is not io.Reader.
type Reader interface {
	Read(data interface{}) error
	Seek(offset int64, whence int) (int64, error)
	UInt8() (data uint8, err error)
	UInt8s(n int) (data []uint8, err error)
	UInt16() (data uint16, err error)
	UInt16s(n int) (data []uint16, err error)
	UInt32() (data uint32, err error)
	UInt32s(n int) (data []uint32, err error)
	UInt64() (data uint64, err error)
	UInt64s(n int) (data []uint64, err error)
	Byte() (data byte, err error)
	Bytes(n int) (data []byte, err error)
	String() (string, error)
	Strings(n int) (data []string, err error)
}
