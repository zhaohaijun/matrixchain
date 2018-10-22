package common

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/ontio/ontology/common/serialization"
)

func BenchmarkZeroCopySource(b *testing.B) {
	const N = 12000
	buf := make([]byte, N)
	rand.Read(buf)

	for i := 0; i < b.N; i++ {
		source := NewZeroCopySource(buf)
		for j := 0; j < N/100; j++ {
			source.NextUint16()
			source.NextByte()
			source.NextUint64()
			source.NextVarUint()
			source.NextBytes(20)
		}
	}

}

func BenchmarkDerserialize(b *testing.B) {
	const N = 12000
	buf := make([]byte, N)
	rand.Read(buf)

	for i := 0; i < b.N; i++ {
		reader := bytes.NewBuffer(buf)
		for j := 0; j < N/100; j++ {
			serialization.ReadUint16(reader)
			serialization.ReadByte(reader)
			serialization.ReadUint64(reader)
			serialization.ReadVarUint(reader, 0)
			serialization.ReadBytes(reader, 20)
		}
	}

}
