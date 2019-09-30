package binutil

import (
	"bytes"
	"encoding/binary"
	"testing"
)

type testWriterSeeker struct {
	buffer *bytes.Buffer
}

func (t *testWriterSeeker) Write(p []byte) (n int, err error) {
	if t.buffer != nil {
		return t.buffer.Write(p)
	}
	// ignore p if t.buffer is nil.
	return len(p), nil
}

func (t *testWriterSeeker) Seek(offset int64, whence int) (int64, error) {
	return offset, nil
}

func TestNewWriter(t *testing.T) {
	testBuffer := &testWriterSeeker{}
	got, ok := NewWriter(testBuffer).(*writerImpl)
	if !ok {
		t.Fatal("got's type is not *binutil.writerImpl as binutil.Writer")
	}
	wont := writerImpl{
		endian:      binary.BigEndian,
		writeSeeker: testBuffer,
	}
	if got == nil {
		t.Fatal("got is nil")
	}
	if got.writeSeeker != wont.writeSeeker {
		t.Error("got.writeSeeker is not equal wont.writeSeeker")
	}
	if got.endian != wont.endian {
		t.Error("got.endian is not equal wont.endian")
	}
}

func TestWriter_String(t *testing.T) {
	testBuffer := &testWriterSeeker{buffer: &bytes.Buffer{}}
	writer := NewWriter(testBuffer)
	data := "Hello, World"
	if err := writer.String(data); err != nil {
		t.Fatal(err)
	}
	wonts := []byte{'H', 'e', 'l', 'l', 'o', ',', ' ', 'W', 'o', 'r', 'l', 'd', 0}
	gots := testBuffer.buffer.Bytes()
	if len(wonts) != len(gots) {
		t.Fatal("len(wonts) != len(gots) :", len(wonts), len(gots))
	}
	for i, got := range gots {
		wb := wonts[i]
		gb := got
		if wb != gb {
			t.Error("wont[", i, "],got[", i, "] = ", wb, gb)
		}
	}
}

func TestWriter_Strings(t *testing.T) {
	testBuffer := &testWriterSeeker{buffer: &bytes.Buffer{}}
	writer := NewWriter(testBuffer)
	data := []string{"Hello", "Good Morning", "Good Evening", "Good Night"}
	if err := writer.Strings(data); err != nil {
		t.Fatal(err)
	}
	wonts := []byte{
		'H', 'e', 'l', 'l', 'o', 0,
		'G', 'o', 'o', 'd', ' ', 'M', 'o', 'r', 'n', 'i', 'n', 'g', 0,
		'G', 'o', 'o', 'd', ' ', 'E', 'v', 'e', 'n', 'i', 'n', 'g', 0,
		'G', 'o', 'o', 'd', ' ', 'N', 'i', 'g', 'h', 't', 0,
	}
	gots := testBuffer.buffer.Bytes()
	if len(wonts) != len(gots) {
		t.Fatal("len(wonts) != len(gots) :", len(wonts), len(gots))
	}
	for i, got := range gots {
		wb := wonts[i]
		gb := got
		if wb != gb {
			t.Error("wont[", i, "],got[", i, "] = ", wb, gb)
		}
	}
}
