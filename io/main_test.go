package main_test

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom
	io.WriterTo
}

func TestCopy(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}

func TestCopyNegative(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello")
	io.Copy(wb, &io.LimitedReader{R: rb, N: -1})
	if wb.String() != "" {
		t.Errorf("Copy on LimitedReader with N<0 copied data")
	}

	io.CopyN(wb, rb, -1)
	if wb.String() != "" {
		t.Errorf("CopyN with N<0 copied data")
	}
}

func TestCopyBuffer(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.CopyBuffer(wb, rb, make([]byte, 1))
	if wb.String() != "hello, world." {
		t.Errorf("CopyBuffer did not work properly")
	}
}

func TestCopyBufferNil(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.CopyBuffer(wb, rb, nil)
	if wb.String() != "hello, world." {
		t.Errorf("CopyBuffer did not work properly")
	}
}

func TestCopyReadFrom(t *testing.T) {
	rb := new(Buffer)
	wb := new(bytes.Buffer) // implements ReadFrom.
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}

func TestCopyWriteTo(t *testing.T) {
	rb := new(bytes.Buffer) // implements WriteTo.
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}

type writeToChecker struct {
	bytes.Buffer
	writeToCalled bool
}

func (wt *writeToChecker) WriteTo(w io.Writer) (int64, error) {
	wt.writeToCalled = true
	return wt.Buffer.WriteTo(w)
}

func TestCopyPriority(t *testing.T) {
	rb := new(writeToChecker)
	wb := new(bytes.Buffer)
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	} else if !rb.writeToCalled {
		t.Errorf("WriteTo was not prioritized over ReadFrom")
	}
}

type zeroErrReader struct {
	err error
}

func (r zeroErrReader) Read(p []byte) (int, error) {
	return copy(p, []byte{0}), r.err
}

type errWriter struct {
	err error
}

func (w errWriter) Write([]byte) (int, error) {
	return 0, w.err
}

func TestCopyReadErrWriterErr(t *testing.T) {
	er, ew := errors.New("readError"), errors.New("writeError")
	r, w := zeroErrReader{err: er}, errWriter{err: ew}
	n, err := io.Copy(w, r)
	if n != 0 || err != ew {
		t.Errorf("Copy(zeroErrReader, errWriter) = %d, %v; want 0, writeError", n, err)
	}
}

func TestCopyN(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.CopyN(wb, rb, 5)
	if wb.String() != "hello" {
		t.Errorf("CopyN did not work properly")
	}
}

func TestCopyNReadFrom(t *testing.T) {
	rb := new(Buffer)
	wb := new(bytes.Buffer)
	rb.WriteString("hello")
	io.CopyN(wb, rb, 5)
	if wb.String() != "hello" {
		t.Errorf("CopyN did not work properly")
	}
}

func TestCopyNWriteTo(t *testing.T) {
	rb := new(bytes.Buffer)
	wb := new(Buffer)
	rb.WriteString("hello")
	io.CopyN(wb, rb, 5)
	if wb.String() != "hello" {
		t.Errorf("CopyN did not work properly")
	}
}

func BenchmarkCopyNSmall(b *testing.B) {
	bs := bytes.Repeat([]byte{0}, 512+1)
	rd := bytes.NewReader(bs)
	buf := new(bytes.Buffer)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.CopyN(buf, rd, 512)
		rd.Reset(bs)
	}
}

func TestTeeReader(t *testing.T) {
	src := []byte("hello, world")
	dst := make([]byte, len(src))
	rb := bytes.NewBuffer(src)
	wb := new(bytes.Buffer)
	r := io.TeeReader(rb, wb)
	if n, err := io.ReadFull(r, dst); err != nil || n != len(src) {
		t.Fatalf("ReadFull(r, dst) = %d, %v; want %d, nil", n, err, len(src))
	}
	if !bytes.Equal(dst, src) {
		t.Errorf("bytes read = %q want %q", dst, src)
	}
	if !bytes.Equal(wb.Bytes(), src) {
		t.Errorf("bytes written = %q want %q", wb.Bytes(), src)
	}
	if n, err := r.Read(dst); n != 0 || err != io.EOF {
		t.Errorf("r.Read at EOF = %d, %v want 0, EOF", n, err)
	}
	rb = bytes.NewBuffer(src)
	pr, pw := io.Pipe()
	pr.Close()
	r = io.TeeReader(rb, pw)
	if n, err := io.ReadFull(r, dst); n != 0 || err != io.ErrClosedPipe {
		t.Errorf("closed tee: ReadFull(r, dst) = %d, %v; want 0, EPIPE", n, err)
	}
}
