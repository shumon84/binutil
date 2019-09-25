package binutil

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
	"testing"
)

type testReaderSeeker struct {
	infill byte
	reader io.Reader
}

func (t *testReaderSeeker) Read(p []byte) (n int, err error) {
	if t.reader != nil {
		return t.reader.Read(p)
	}
	for i := range p {
		p[i] = t.infill
	}
	return len(p), nil
}

func (t *testReaderSeeker) Seek(offset int64, whence int) (int64, error) {
	return offset, nil
}

func TestNewReader(t *testing.T) {
	testBuffer := &testReaderSeeker{}
	got, ok := NewReader(testBuffer).(*readerImpl)
	if !ok {
		t.Fatal("got's type is not *binutil.readerImple as binutil.Reader")
	}
	wont := readerImpl{
		endian:     binary.BigEndian,
		readSeeker: testBuffer,
	}
	if got == nil {
		t.Fatal("got is nil")
	}
	if got.readSeeker != wont.readSeeker {
		t.Error("got.readSeeker is not equal wont.readSeeker")
	}
	if got.endian != wont.endian {
		t.Error("got.endian is not equal wont.endian")
	}
}

func TestReader_UInt8(t *testing.T) {
	var wont uint8 = 0x01
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	got, err := reader.UInt8()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error()
	}
}

func TestReader_UInt8s(t *testing.T) {
	var wont uint8 = 0x01
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	gots, err := reader.UInt8s(10)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		if got != wont {
			t.Error("got[", i, "] != wont")
		}
	}
}

func TestReader_UInt16(t *testing.T) {
	var wont uint16 = 0x0101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	got, err := reader.UInt16()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error()
	}
}

func TestReader_UInt16s(t *testing.T) {
	var wont uint16 = 0x0101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	gots, err := reader.UInt16s(10)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		if got != wont {
			t.Error("got[", i, "] != wont")
		}
	}
}

func TestReader_UInt32(t *testing.T) {
	var wont uint32 = 0x01010101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	got, err := reader.UInt32()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error()
	}
}

func TestReader_UInt32s(t *testing.T) {
	var wont uint32 = 0x01010101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	gots, err := reader.UInt32s(10)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		if got != wont {
			t.Error("got[", i, "] != wont")
		}
	}
}

func TestReader_UInt64(t *testing.T) {
	var wont uint64 = 0x0101010101010101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	got, err := reader.UInt64()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error()
	}
}

func TestReader_UInt64s(t *testing.T) {
	var wont uint64 = 0x0101010101010101
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	gots, err := reader.UInt64s(10)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		if got != wont {
			t.Error("got[", i, "] != wont")
		}
	}
}

func TestReader_Byte(t *testing.T) {
	var wont byte = 0x01
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	got, err := reader.Byte()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error("got != wont")
	}
}

func TestReader_Bytes(t *testing.T) {
	var wont byte = 0x01
	testBuffer := &testReaderSeeker{infill: 0x01}
	reader := NewReader(testBuffer)
	gots, err := reader.Bytes(10)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range gots {
		if got != wont {
			t.Error("got[", i, "] != wont")
		}
	}
}

func TestReader_String(t *testing.T) {
	wont := "Hello, World"
	testBuffer := &testReaderSeeker{reader: strings.NewReader(wont)}
	reader := NewReader(testBuffer)
	got, err := reader.String()
	if err != nil {
		t.Fatal(err)
	}
	if got != wont {
		t.Error("got != wont")
	}
}

func TestReader_Strings(t *testing.T) {
	wont := []string{"Hello", "Good Morning", "Good Evening", "Good Night"}
	data := []byte{
		'H', 'e', 'l', 'l', 'o', 0,
		'G', 'o', 'o', 'd', ' ', 'M', 'o', 'r', 'n', 'i', 'n', 'g', 0,
		'G', 'o', 'o', 'd', ' ', 'E', 'v', 'e', 'n', 'i', 'n', 'g', 0,
		'G', 'o', 'o', 'd', ' ', 'N', 'i', 'g', 'h', 't', 0,
	}

	testBuffer := &testReaderSeeker{reader: bytes.NewBuffer(data)}
	reader := NewReader(testBuffer)

	gots, err := reader.Strings(len(wont))
	if err != nil {
		t.Fatal(err)
	}
	if len(gots) != len(wont) {
		t.Error("len(gots) =", len(gots))
	}
	for i, got := range gots {
		wb := wont[i]
		gb := got
		for i := range gb {
			if gb[i] != wb[i] {
				t.Error("w[", i, "],g[", i, "] = ", wb[i], gb[i])
			}
		}
	}
}
