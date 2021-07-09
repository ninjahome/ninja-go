package websocket

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	lvEndDelim string = "@@@@"
	lvHeadLen  int    = 4
)

type LVReaderWriter struct {
	io.ReadWriter
	buflen int
}

func NewLVRW(rw io.ReadWriter, buflen int) *LVReaderWriter {
	return &LVReaderWriter{
		ReadWriter: rw,
		buflen:     buflen,
	}
}

func (lv *LVReaderWriter) Read(p []byte) (n int, err error) {

	buf := make([]byte, lvHeadLen)
	n, err = lv.ReadWriter.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != lvHeadLen {
		return 0, errors.New("read error")
	}

	nl := binary.BigEndian.Uint32(buf)
	if int(nl) > lv.buflen || int(nl) > len(p) {
		err = fmt.Errorf("length not correct, nl: %d, len(p): %d", nl, len(p))
		return 0, err
	}

	n, err = lv.ReadWriter.Read(p)
	if err != nil {
		return 0, err
	}
	if n != int(nl) {
		err = fmt.Errorf("read not correct, nl: %d, n: %d", nl, n)
		return 0, err
	}

	return n, err
}

func IsReadEnd(p []byte) bool {
	if len(p) == len(lvEndDelim) {
		if string(p) == lvEndDelim {
			return true
		}
	}
	return false
}

func (lv *LVReaderWriter) Write(p []byte) (n int, err error) {
	np := len(p)
	if np > lv.buflen {
		err = fmt.Errorf("write check error,np:%d, buflen: %d", np, lv.buflen)
		return 0, err
	}

	buf := make([]byte, lvHeadLen)
	binary.BigEndian.PutUint32(buf, uint32(np))
	n, err = lv.ReadWriter.Write(buf)
	if err != nil {
		return 0, err
	}
	if n != lvHeadLen {
		err = fmt.Errorf("write head not correct, n: %d", n)
		return 0, err
	}

	n, err = lv.ReadWriter.Write(p)
	if err != nil {
		return 0, err
	}
	if n != np {
		err = fmt.Errorf("write data not correct, n: %d,np: %d", n, np)
		return 0, err
	}

	return n, nil
}

func (lv *LVReaderWriter) Commit() (n int, err error) {
	buf := make([]byte, lvHeadLen)
	binary.BigEndian.PutUint32(buf, uint32(lvHeadLen))

	data := make([]byte, lvHeadLen+len(lvEndDelim))
	copy(data, buf)
	copy(data[lvHeadLen:], []byte(lvEndDelim))

	n, err = lv.ReadWriter.Write(data)
	if err != nil {
		return 0, err
	}
	if n != lvHeadLen+len(lvEndDelim) {
		err = fmt.Errorf("commit head not correct, n: %d", n)
		return 0, err
	}

	return n, nil
}

func (lv *LVReaderWriter) ReadFull(buf []byte) (int, error) {
	nb := len(buf)
	totalLen := 0
	for {
		tmpbuf := make([]byte, lv.buflen)
		n, err := lv.Read(tmpbuf)
		if err != nil {
			return 0, err
		}
		if IsReadEnd(tmpbuf[:n]) {
			return totalLen, nil
		}

		if nb < totalLen+n {
			return 0, errors.New("buf not enough")
		}

		copy(buf[totalLen:], tmpbuf[:n])
		totalLen += n
	}
}
