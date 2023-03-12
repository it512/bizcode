package bizcode

import "io"

func nonce(buf []byte) {
	io.ReadFull(src, buf)

	for i, b := range buf {
		buf[i] = chars[b%62]
	}
}

func nonce2(n int) []byte {
	b := make([]byte, n)
	nonce(b)
	return b
}
