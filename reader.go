package binutil

import (
	"encoding/binary"
	"io"
)

type readerImpl struct {
	readSeeker io.ReadSeeker
	endian     binary.ByteOrder
}

func NewReader(readSeeker io.ReadSeeker) Reader {
	return NewReaderWithEndian(readSeeker, binary.BigEndian)
}

func NewReaderWithEndian(readSeeker io.ReadSeeker, byteOrder binary.ByteOrder) Reader {
	return &readerImpl{
		readSeeker: readSeeker,
		endian:     byteOrder,
	}
}

func (r *readerImpl) Read(data interface{}) error {
	return binary.Read(r.readSeeker, r.endian, data)
}

func (r *readerImpl) Seek(offset int64, whence int) (int64, error) {
	return r.readSeeker.Seek(offset, whence)
}

func (r *readerImpl) UInt8() (data uint8, err error) {
	err = binary.Read(r.readSeeker, r.endian, &data)
	return
}

func (r *readerImpl) UInt8s(n int) (data []uint8, err error) {
	data = make([]uint8, n)
	err = binary.Read(r.readSeeker, r.endian, data)
	return
}

func (r *readerImpl) UInt16() (data uint16, err error) {
	err = binary.Read(r.readSeeker, r.endian, &data)
	return
}

func (r *readerImpl) UInt16s(n int) (data []uint16, err error) {
	data = make([]uint16, n)
	err = binary.Read(r.readSeeker, r.endian, data)
	return
}

func (r *readerImpl) UInt32() (data uint32, err error) {
	err = binary.Read(r.readSeeker, r.endian, &data)
	return
}

func (r *readerImpl) UInt32s(n int) (data []uint32, err error) {
	data = make([]uint32, n)
	err = binary.Read(r.readSeeker, r.endian, data)
	return
}

func (r *readerImpl) UInt64() (data uint64, err error) {
	err = binary.Read(r.readSeeker, r.endian, &data)
	return
}

func (r *readerImpl) UInt64s(n int) (data []uint64, err error) {
	data = make([]uint64, n)
	err = binary.Read(r.readSeeker, r.endian, data)
	return
}

func (r *readerImpl) Byte() (data byte, err error) {
	err = binary.Read(r.readSeeker, r.endian, &data)
	return
}

func (r *readerImpl) Bytes(n int) (data []byte, err error) {
	data = make([]byte, n)
	err = binary.Read(r.readSeeker, r.endian, data)
	return
}

func (r *readerImpl) String() (string, error) {
	str := make([]byte, 0)
	for {
		c := byte(0)
		err := binary.Read(r.readSeeker, r.endian, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if c == 0 {
			break
		}
		str = append(str, c)
	}
	return string(str), nil
}

func (r *readerImpl) Strings(n int) (data []string, err error) {
	data = make([]string, n)
	for i := range data {
		data[i], err = r.String()
		if err != nil {
			return nil, err
		}
	}
	return
}
