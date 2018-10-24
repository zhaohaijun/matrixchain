package utils

import (
	"bytes"
	"encoding/binary"
)

type VmReader struct {
	reader     *bytes.Reader
	BaseStream []byte
}

func NewVmReader(b []byte) *VmReader {
	var vmreader VmReader
	vmreader.reader = bytes.NewReader(b)
	vmreader.BaseStream = b
	return &vmreader
}
func (r *VmReader) ReadByte() (byte, error) {
	byte, err := r.reader.ReadByte()
	return byte, err
}
func (r *VmReader) ReadBytes(count int) []byte {
	b := make([]byte, count)
	r.reader.Read(b)
	return b
}
func (r *VmReader) ReadBytesInto(b []byte) {
	r.reader.Read(b)
}
func (r *VmReader) ReadUint16() uint16 {
	var b [2]byte
	r.ReadBytesInto(b[:])
	return binary.LittleEndian.Uint16(b[:])
}
func (r *VmReader) ReadUint32() uint32 {
	var b [4]byte
	r.ReadBytesInto(b[:])
	return binary.LittleEndian.Uint32(b[:])
}

func (r *VmReader) ReadUint64() uint64 {
	var b [8]byte
	r.ReadBytesInto(b[:])
	return binary.LittleEndian.Uint64(b[:])
}

func (r *VmReader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

func (r *VmReader) ReadInt32() int32 {
	return int32(r.ReadUint32())
}

func (r *VmReader) Position() int {
	return int(r.reader.Size()) - r.reader.Len()
}

func (r *VmReader) Length() int {
	return r.reader.Len()
}

func (r *VmReader) Seek(offset int64, whence int) (int64, error) {
	return r.reader.Seek(offset, whence)
}

func (r *VmReader) ReadVarBytes(max uint32) []byte {
	n := int(r.ReadVarInt(uint64(max)))
	return r.ReadBytes(n)
}

func (r *VmReader) ReadVarInt(max uint64) uint64 {
	fb, _ := r.ReadByte()
	var value uint64

	switch fb {
	case 0xFD:
		value = uint64(r.ReadInt16())
	case 0xFE:
		value = uint64(r.ReadUint32())
	case 0xFF:
		value = uint64(r.ReadUint64())
	default:
		value = uint64(fb)
	}
	if value > max {
		return 0
	}
	return value
}

func (r *VmReader) ReadVarString(maxlen uint32) string {
	bs := r.ReadVarBytes(maxlen)
	return string(bs)
}
