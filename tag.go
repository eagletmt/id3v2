package id3v2

import (
	"io"
)

type Tag struct {
	Header         Header
	Frames         map[string]*Frame
	UserTextFrames []*UserTextFrame
}

func (t *Tag) Read(r io.Reader) error {
	err := t.Header.Read(r)
	if err != nil {
		return err
	}

	t.Frames = make(map[string]*Frame)
	t.UserTextFrames = make([]*UserTextFrame, 0)
	for size := t.Header.Size; size > 0; {
		frame := &Frame{}
		err = frame.Read(r)
		if err != nil {
			break
		}
		if frame.FrameId == "TXXX" {
			t.UserTextFrames = append(t.UserTextFrames, decodeUserTextFrame(frame))
		} else {
			t.Frames[frame.FrameId] = frame
		}
		size -= 10 + frame.Size
	}
	return nil
}
