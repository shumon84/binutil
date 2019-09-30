package binutil

import (
	"encoding/binary"
	"io"
)

type writerImpl struct {
	writeSeeker io.WriteSeeker
	endian      binary.ByteOrder
}

func NewWriter(writeSeeker io.WriteSeeker) Writer {
	return NewWriterWithEndian(writeSeeker, binary.BigEndian)
}

func NewWriterWithEndian(writeSeeker io.WriteSeeker, byteOrder binary.ByteOrder) Writer {
	return &writerImpl{
		writeSeeker: writeSeeker,
		endian:      byteOrder,
	}
}

func (w *writerImpl) Write(data interface{}) error {
	return binary.Write(w.writeSeeker, w.endian, data)
}

func (w *writerImpl) Seek(offset int64, whence int) (int64, error) {
	return w.writeSeeker.Seek(offset, whence)
}

func (w *writerImpl) UInt8(data uint8) error {
	return w.Write(data)
}

func (w *writerImpl) UInt8s(data []uint8) error {
	return binary.Write(w.writeSeeker, w.endian, data)
}

func (w *writerImpl) UInt16(data uint16) error {
	return binary.Write(w.writeSeeker, w.endian, data)
}

func (w *writerImpl) UInt16s(data []uint16) error {
	return w.Write(data)
}

func (w *writerImpl) UInt32(data uint32) error {
	return w.Write(data)
}

func (w *writerImpl) UInt32s(data []uint32) error {
	return w.Write(data)
}

func (w *writerImpl) UInt64(data uint64) error {
	return w.Write(data)
}

func (w *writerImpl) UInt64s(data []uint64) error {
	return w.Write(data)
}

func (w *writerImpl) Byte(data byte) error {
	return w.Write(data)
}

func (w *writerImpl) Bytes(data []byte) error {
	return w.Write(data)
}

func (w *writerImpl) String(data string) error {
	if err := w.Write([]byte(data)); err != nil {
		return err
	}
	if err := w.Write(byte(0)); err != nil {
		return err
	}
	return nil
}

func (w *writerImpl) Strings(data []string) error {
	for _, d := range data {
		if err := w.String(d); err != nil {
			return err
		}
	}
	return nil
}
