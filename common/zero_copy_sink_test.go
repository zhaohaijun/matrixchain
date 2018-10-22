package common

import (
	"testing"

	"bytes"

	ser "github.com/ontio/ontology/common/serialization"
	"github.com/stretchr/testify/assert"
)

func TestSourceSink(t *testing.T) {
	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	sink := NewZeroCopySink(nil)
	sink.WriteByte(a3)
	sink.WriteUint16(a4)
	sink.WriteUint32(a5)
	sink.WriteUint64(a6)
	sink.WriteVarUint(a7)
	sink.WriteVarBytes(a8)
	sink.WriteString(a9)

	source := NewZeroCopySource(sink.Bytes())
	b3, _ := source.NextByte()
	assert.Equal(t, a3, b3)
	b4, _ := source.NextUint16()
	assert.Equal(t, a4, b4)
	b5, _ := source.NextUint32()
	assert.Equal(t, a5, b5)
	b6, _ := source.NextUint64()
	assert.Equal(t, a6, b6)
	b7, _, _, _ := source.NextVarUint()
	assert.Equal(t, a7, b7)
	b8, _, _, _ := source.NextVarBytes()
	assert.Equal(t, a8, b8)
	b9, _, _, _ := source.NextString()
	assert.Equal(t, a9, b9)

}

func BenchmarkSerialize(ben *testing.B) {
	N := 1000
	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	b := new(bytes.Buffer)
	for i := 0; i < ben.N; i++ {
		b.Reset()
		for j := 0; j < N; j++ {
			ser.WriteVarUint(b, uint64(a3))
			ser.WriteVarUint(b, uint64(a4))
			ser.WriteVarUint(b, uint64(a5))
			ser.WriteVarUint(b, uint64(a6))
			ser.WriteVarUint(b, uint64(a7))
			ser.WriteVarBytes(b, a8)
			ser.WriteString(b, a9)

			b.WriteByte(20)
			b.WriteByte(21)
			b.WriteByte(22)
		}
	}
}

func BenchmarkZeroCopySink(ben *testing.B) {
	N := 1000
	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	sink := NewZeroCopySink(nil)
	for i := 0; i < ben.N; i++ {
		sink.Reset()
		for j := 0; j < N; j++ {
			sink.WriteVarUint(uint64(a3))
			sink.WriteVarUint(uint64(a4))
			sink.WriteVarUint(uint64(a5))
			sink.WriteVarUint(uint64(a6))
			sink.WriteVarUint(uint64(a7))
			sink.WriteVarBytes(a8)
			sink.WriteString(a9)
			sink.WriteByte(20)
			sink.WriteByte(21)
			sink.WriteByte(22)
		}
	}

}
