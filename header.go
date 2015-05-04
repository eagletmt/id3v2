package id3v2

import (
	"fmt"
	"io"
)

type Header struct {
	Version Version
	Flags   Flags
	Size    uint64
}

type Version struct {
	Major    uint8
	Revision uint8
}

type Flags struct {
	Unsynchronisation     bool
	ExtendedHeader        bool
	ExperimentalIndicator bool
	FooterPresent         bool
}

func (h *Header) Read(r io.Reader) error {
	err := readFileIdentifier(r)
	if err != nil {
		return err
	}
	h.Version, err = readVersion(r)
	if err != nil {
		return err
	}
	err = h.readFlags(r)
	if err != nil {
		return err
	}
	h.Size, err = readSize(r)
	if err != nil {
		return err
	}
	return nil
}

func readFileIdentifier(r io.Reader) error {
	buf := make([]byte, 3)
	err := readBytes(buf, r, 3)
	if err != nil {
		return err
	}
	if buf[0] == 'I' && buf[1] == 'D' && buf[2] == '3' {
		return nil
	} else {
		return fmt.Errorf("Not ID3v2")
	}
}

func readVersion(r io.Reader) (Version, error) {
	buf := make([]byte, 2)
	err := readBytes(buf, r, 2)
	if err != nil {
		return Version{}, err
	}
	return Version{Major: uint8(buf[0]), Revision: uint8(buf[1])}, nil
}

func (h *Header) readFlags(r io.Reader) error {
	buf := make([]byte, 1)
	err := readBytes(buf, r, 1)
	if err != nil {
		return err
	}

	h.Flags.Unsynchronisation = (buf[0] & (1 << 0)) != 0     // 3.1.a
	h.Flags.ExtendedHeader = (buf[0] & (1 << 1)) != 0        // 3.1.b
	h.Flags.ExperimentalIndicator = (buf[0] & (1 << 2)) != 0 // 3.1.c
	h.Flags.FooterPresent = (buf[0] & (1 << 3)) != 0         // 3.1.d

	return nil
}

func readSize(r io.Reader) (uint64, error) {
	buf := make([]byte, 4)
	err := readBytes(buf, r, 4)
	if err != nil {
		return 0, err
	}

	return synchSafe(buf)
}
