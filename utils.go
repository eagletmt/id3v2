package id3v2

import (
	"fmt"
	"io"
)

func readBytes(buf []byte, r io.Reader, n int) error {
	m, err := r.Read(buf)
	if err != nil {
		return err
	}
	if m != n {
		return fmt.Errorf("read only %d bytes, expected %d bytes", m, n)
	}
	return nil
}

func synchSafe(buf []byte) (uint64, error) {
	// 6.2 Synchsafe integers
	n := uint64(0)
	for i, b := range buf {
		if (b & (1 << 7)) != 0 {
			return 0, fmt.Errorf("Invalid synchsafe integer found at %d: %d", i, b)
		}
		n |= (n << 7) | uint64(b)
	}
	return n, nil
}

func toUint64(buf []byte) uint64 {
	return uint64(buf[0])<<24 | uint64(buf[1])<<16 | uint64(buf[2])<<8 | uint64(buf[3])
}
