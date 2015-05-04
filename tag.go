package id3v2

import (
	"io"
)

type Tag struct {
	Header Header
	Frames map[string]*Frame
}

func (t *Tag) Read(r io.Reader) error {
	err := t.Header.Read(r)
	if err != nil {
		return err
	}

	t.Frames = make(map[string]*Frame)
	for size := t.Header.Size; size > 0; {
		frame := &Frame{}
		err = frame.Read(r)
		if err != nil {
			break
		}
		t.Frames[frame.FrameId] = frame
		size -= 10 + frame.Size
	}
	return nil
}
