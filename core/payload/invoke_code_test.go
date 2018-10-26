package payload

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvokeCode_Serialize(t *testing.T) {
	code := InvokeCode{
		Code: []byte{1, 2, 3},
	}

	buf := bytes.NewBuffer(nil)
	code.Serialize(buf)
	bs := buf.Bytes()
	var code2 InvokeCode
	code2.Deserialize(buf)
	assert.Equal(t, code, code2)

	buf = bytes.NewBuffer(bs[:len(bs)-2])
	err := code.Deserialize(buf)

	assert.NotNil(t, err)
}
