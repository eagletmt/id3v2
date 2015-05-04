package id3v2

import (
	"fmt"
	"io"
)

type Frame struct {
	FrameId     string
	Size        uint64
	StatusFlags FrameStatusFlags
	FormatFlags FrameFormatFlags
	Payload     []byte
}

type FrameStatusFlags struct {
	TagAlterPreservation  bool
	FileAlterPreservation bool
	ReadOnly              bool
}

type FrameFormatFlags struct {
	GroupingIdentity    bool
	Compression         bool
	Encryption          bool
	Unsynchronisation   bool
	DataLengthIndicator bool
}

// 4
func (f *Frame) Read(r io.Reader) error {
	buf := make([]byte, 4)
	err := readBytes(buf, r, 4)
	if err != nil {
		return err
	}
	f.FrameId = string(buf)
	err = readBytes(buf, r, 4)
	if err != nil {
		return err
	}
	f.Size = toUint64(buf)
	if f.Size == 0 {
		return fmt.Errorf("Empty frame")
	}
	err = f.readFlags(r)
	if err != nil {
		return err
	}
	f.Payload = make([]byte, f.Size)
	err = readBytes(f.Payload, r, len(f.Payload))
	if err != nil {
		return err
	}
	return nil
}

// 4.1
func (f *Frame) readFlags(r io.Reader) error {
	buf := make([]byte, 2)
	err := readBytes(buf, r, 2)
	if err != nil {
		return err
	}

	f.StatusFlags.read(buf[0])
	f.FormatFlags.read(buf[1])
	return nil
}

// 4.1.1
func (f *FrameStatusFlags) read(b byte) {
	f.TagAlterPreservation = (b & (1 << 6)) != 0
	f.FileAlterPreservation = (b & (1 << 5)) != 0
	f.ReadOnly = (b & 1 << 4) != 0
}

// 4.1.2
func (f *FrameFormatFlags) read(b byte) {
	f.GroupingIdentity = (b & (1 << 6)) != 0
	f.Compression = (b & (1 << 3)) != 0
	f.Encryption = (b & (1 << 2)) != 0
	f.Unsynchronisation = (b & (1 << 1)) != 0
	f.DataLengthIndicator = (b & (1 << 0)) != 0
}
