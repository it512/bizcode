package bizcode

import (
	"bufio"
	"bytes"
	crand "crypto/rand"
	"errors"
	"io"
	"sync"
	"time"
	"unsafe"
)

type entropyPool struct {
	mux    sync.Mutex
	buffer io.Reader
}

func newEntropyPool() *entropyPool {
	return &entropyPool{
		buffer: bufio.NewReader(crand.Reader),
	}
}

func (r *entropyPool) Read(p []byte) (n int, err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.buffer.Read(p)
}

var src = newEntropyPool()

const layout = "20060102150405"

var (
	z = []byte{'0', '0'}
)

func New() string {
	return NewBy(z)
}

func NewBy(typ []byte) string {
	return str(build(make([]byte, 0, 32), typ))
}

func str(c []byte) string {
	if len(c) != 32 {
		panic("len != 32")
	}
	return unsafe.String(&c[0], 32)
}

func build(c []byte, typ []byte) []byte {

	t := time.Now()
	b := t.AppendFormat(c[:0], layout)

	if len(typ) >= 2 {
		b = append(b, typ[0], typ[1])
	} else {
		b = append(b, z...)
	} // 14 +2 = 16

	b = append(b, nonce2(12)...)
	b = append(b, x(b))
	b = append(b, b62(crc16(b))...)
	return b

}

func By(typ string) func() string {
	if len(typ) >= 2 {
		b := []byte{typ[0], typ[1]}
		return func() string { return NewBy(b) }
	}
	return func() string { return NewBy(z) }
}

var (
	ErrCodeLen = errors.New("Len != 32")
	ErrCodeSum = errors.New("CheckSum error")
)

func CheckCode(code string, funcs ...func([]byte) error) error {
	if len(code) != 32 {
		return ErrCodeLen
	}

	if len(funcs) == 0 {
		return nil
	}

	bs := unsafe.Slice(unsafe.StringData(code), 32)
	for _, f := range funcs {
		if err := f(bytes.Clone(bs)); err != nil {
			return err
		}
	}

	return nil
}

func x(bs []byte) byte {
	i := 0
	for _, b := range bs {
		if b >= '0' && b <= '9' {
			i++
		}
	}
	return chars[i-(28-i)]
}

func CodeType(code string) string {
	return code[14:16]
}
