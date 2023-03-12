package bizcode

import (
	"testing"
)

// 202303122248210AaqMrIU33acws60Jd
func Test_b62_0(t *testing.T) {
	b := []byte("202303122248210AaqMrIU33acws6")
	i := crc16(b)
	r := b62(i)
	if len(r) != 3 && r[0] != '0' {
		t.Fail()
	}

}
