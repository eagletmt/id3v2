package id3v2

import (
	"bytes"
	"log"
)

// 4.14

type AttachedPictureFrame struct {
	MimeType    string
	PictureType uint8
	Description string
	PictureData []byte
}

func (t *Tag) AttachedPicture() *AttachedPictureFrame {
	frame, ok := t.Frames["APIC"]
	if !ok {
		return nil
	}
	return decodeAttachedPictureFrame(frame.Payload)
}

func decodeAttachedPictureFrame(payload []byte) *AttachedPictureFrame {
	encoding := payload[0]
	i := bytes.IndexByte(payload[1:], 0x00)
	if i == -1 {
		log.Printf("Cannot find MIME type termination in APIC frame")
		return nil
	}
	f := &AttachedPictureFrame{}
	f.MimeType = string(payload[1 : i+1])
	f.PictureType = payload[i+2]
	rest := payload[i+3:]
	j := bytes.IndexByte(rest, 0x00)
	if j == -1 {
		log.Printf("Cannot find Description termination in APIc frame")
		return nil
	}
	f.Description = decodeText(encoding, rest[:j+1])
	f.PictureData = rest[j+1:]
	return f
}
