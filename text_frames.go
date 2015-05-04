package id3v2

import (
	"log"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// 4.2.1

func (t *Tag) ContentGroup() string {
	return t.getTextInformationFrame("TIT1")
}

func (t *Tag) Title() string {
	return t.getTextInformationFrame("TIT2")
}

func (t *Tag) Subtitle() string {
	return t.getTextInformationFrame("TIT3")
}

func (t *Tag) Album() string {
	return t.getTextInformationFrame("TALB")
}

func (t *Tag) OriginalAlbum() string {
	return t.getTextInformationFrame("TOAL")
}

func (t *Tag) Track() string {
	return t.getTextInformationFrame("TRCK")
}

func (t *Tag) PartOfSet() string {
	return t.getTextInformationFrame("TPOS")
}

func (t *Tag) SetSubtitle() string {
	return t.getTextInformationFrame("TSST")
}

func (t *Tag) ISRC() string {
	return t.getTextInformationFrame("TSRC")
}

// 4.2.2

func (t *Tag) Artist() string {
	return t.getTextInformationFrame("TPE1")
}

func (t *Tag) AlbumArtist() string {
	return t.getTextInformationFrame("TPE2")
}

func (t *Tag) Conductor() string {
	return t.getTextInformationFrame("TPE3")
}

func (t *Tag) InterpretedBy() string {
	return t.getTextInformationFrame("TPE4")
}

func (t *Tag) OriginalArtist() string {
	return t.getTextInformationFrame("TOPE")
}

func (t *Tag) Lyricist() string {
	return t.getTextInformationFrame("TEXT")
}

func (t *Tag) OriginalLyricist() string {
	return t.getTextInformationFrame("TOLY")
}

func (t *Tag) Composer() string {
	return t.getTextInformationFrame("TCOM")
}

func (t *Tag) MusicianCredits() string {
	return t.getTextInformationFrame("TMCL")
}

func (t *Tag) InvolvedPeople() string {
	return t.getTextInformationFrame("TIPL")
}

func (t *Tag) EncodedBy() string {
	return t.getTextInformationFrame("TENC")
}

// 4.2.3

func (t *Tag) BPM() string {
	return t.getTextInformationFrame("TBPM")
}

func (t *Tag) Length() string {
	return t.getTextInformationFrame("TLEN")
}

func (t *Tag) InitialKey() string {
	return t.getTextInformationFrame("TKEY")
}

func (t *Tag) Language() string {
	return t.getTextInformationFrame("TLAN")
}

func (t *Tag) ContentType() string {
	return t.getTextInformationFrame("TCON")
}

func (t *Tag) FileType() string {
	return t.getTextInformationFrame("TFLT")
}

func (t *Tag) MediaType() string {
	return t.getTextInformationFrame("TMED")
}

func (t *Tag) Mood() string {
	return t.getTextInformationFrame("TMOO")
}

// 4.2.4

func (t *Tag) CopyrightMessage() string {
	return t.getTextInformationFrame("TCOP")
}

func (t *Tag) ProducedNotice() string {
	return t.getTextInformationFrame("TPRO")
}

func (t *Tag) Publisher() string {
	return t.getTextInformationFrame("TPUB")
}

func (t *Tag) Licensee() string {
	return t.getTextInformationFrame("TOWN")
}

func (t *Tag) InternetRadioStationName() string {
	return t.getTextInformationFrame("TRSN")
}

func (t *Tag) InternetRadioStationOwner() string {
	return t.getTextInformationFrame("TRSO")
}

// 4.2.5

func (t *Tag) OriginalFilename() string {
	return t.getTextInformationFrame("TOFN")
}

func (t *Tag) PlaylistDelay() string {
	return t.getTextInformationFrame("TDLY")
}

func (t *Tag) EncodingTime() string {
	return t.getTextInformationFrame("TDEN")
}

func (t *Tag) OriginalReleaseTime() string {
	return t.getTextInformationFrame("TDOR")
}

func (t *Tag) RecordingTime() string {
	return t.getTextInformationFrame("TDRC")
}

func (t *Tag) ReleaseTime() string {
	return t.getTextInformationFrame("TDRL")
}

func (t *Tag) TaggingTime() string {
	return t.getTextInformationFrame("TDTG")
}

func (t *Tag) EncodingSettings() string {
	return t.getTextInformationFrame("TSSE")
}

func (t *Tag) AlbumSortOrder() string {
	return t.getTextInformationFrame("TSOA")
}

func (t *Tag) PerformerSortOrder() string {
	return t.getTextInformationFrame("TSOP")
}

func (t *Tag) TitleSortOrder() string {
	return t.getTextInformationFrame("TSOT")
}

func (t *Tag) getTextInformationFrame(frameId string) string {
	frame, ok := t.Frames[frameId]
	if !ok {
		return ""
	}
	return decodeTextInformation(frame.Payload)
}

func decodeTextInformation(buf []byte) string {
	encoding := buf[0]
	switch encoding {
	case 0x00:
		// ISO-8859-1 string terminated with NUL.
		return decodeLatin1(buf[1:])
	case 0x01:
		// UTF-16 string with BOM terminated with NUL.
		return decodeUtf16(buf[1:])
	case 0x02:
		// UTF-16BE string without BOM terminated with NUL.
		return decodeUtf16be(buf[1:])
	case 0x03:
		// UTF-8 string terminated with NUL.
		return string(buf[1:])
	default:
		log.Printf("Unknown encoding: %#x", encoding)
		return ""
	}
}

func decodeLatin1(buf []byte) string {
	utf8 := make([]rune, len(buf))
	for i, b := range buf {
		utf8[i] = rune(b)
	}
	return string(buf)
}

func decodeUtf16(buf []byte) string {
	return decodeUtf16With(buf, unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM))
}

func decodeUtf16be(buf []byte) string {
	return decodeUtf16With(buf, unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM))
}

func decodeUtf16With(buf []byte, enc encoding.Encoding) string {
	decoder := enc.NewDecoder()
	dst := make([]byte, len(buf)*2)
	nDst, _, err := decoder.Transform(dst, buf, true)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(dst[:nDst])
}
