package id3v2

import (
	"io"
)

func Read(r io.Reader) (*Tag, error) {
	tag := &Tag{}
	err := tag.Read(r)
	return tag, err
}
