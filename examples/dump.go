package main

import (
	"fmt"
	"log"
	"os"

	".."
)

func main() {
	exitcode := 0
	for _, path := range os.Args[1:] {
		err := dump(path)
		if err != nil {
			log.Printf("%s: %s", path, err)
			exitcode++
		}
	}
	os.Exit(exitcode)
}

func dump(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tag, err := id3v2.Read(file)
	if err != nil {
		return err
	}

	dumpInfo(path, tag)
	return nil
}

func dumpInfo(path string, tag *id3v2.Tag) {
	fmt.Printf("%s:\n", path)

	fmt.Printf("  Frame IDs:")
	for k, _ := range tag.Frames {
		fmt.Printf(" %s", k)
	}
	fmt.Println()
	fmt.Println()

	fmt.Printf("  Content group: %s\n", tag.ContentGroup())
	fmt.Printf("  Title/Songname/Content: %s\n", tag.Title())
	fmt.Printf("  Subtitle: %s\n", tag.Subtitle())
	fmt.Printf("  Album: %s\n", tag.Album())
	fmt.Printf("  Original album: %s\n", tag.OriginalAlbum())
	fmt.Printf("  Track number: %s\n", tag.Track())
	fmt.Printf("  Part of a set: %s\n", tag.PartOfSet())
	fmt.Printf("  Set subtitle: %s\n", tag.SetSubtitle())
	fmt.Printf("  ISRC: %s\n", tag.ISRC())
	fmt.Println()
	fmt.Printf("  Lead artist/Lead performer/Soloist/Performing group: %s\n", tag.Artist())
	fmt.Printf("  Band/Orchestra/Accompaniment: %s\n", tag.AlbumArtist())
	fmt.Printf("  Conductor: %s\n", tag.Conductor())
	fmt.Printf("  Interpreted, remixed, or otherwise modified by: %s\n", tag.InterpretedBy())
	fmt.Printf("  Original artist/performer: %s\n", tag.OriginalArtist())
	fmt.Printf("  Lyricist/Text writer: %s\n", tag.Lyricist())
	fmt.Printf("  Original lyricist/text writer: %s\n", tag.OriginalLyricist())
	fmt.Printf("  Composer: %s\n", tag.Composer())
	fmt.Printf("  Musician credits list: %s\n", tag.MusicianCredits())
	fmt.Printf("  Involved people list: %s\n", tag.InvolvedPeople())
	fmt.Printf("  Encoded by: %s\n", tag.EncodedBy())
	fmt.Println()
	fmt.Printf("  BPM: %s\n", tag.BPM())
	fmt.Printf("  Length: %s\n", tag.Length())
	fmt.Printf("  Initial key: %s\n", tag.InitialKey())
	fmt.Printf("  Language: %s\n", tag.Language())
	fmt.Printf("  Content type: %s\n", tag.ContentType())
	fmt.Printf("  File type: %s\n", tag.FileType())
	fmt.Printf("  Media type: %s\n", tag.MediaType())
	fmt.Printf("  Mood: %s\n", tag.Mood())
	fmt.Println()
	fmt.Printf("  Copyright message: %s\n", tag.CopyrightMessage())
	fmt.Printf("  Produced notice: %s\n", tag.ProducedNotice())
	fmt.Printf("  Publisher: %s\n", tag.Publisher())
	fmt.Printf("  File owner/licensee: %s\n", tag.Licensee())
	fmt.Printf("  Internet radio station name: %s\n", tag.InternetRadioStationName())
	fmt.Printf("  Internet radio station owner: %s\n", tag.InternetRadioStationOwner())
	fmt.Println()
	fmt.Printf("  Original filename: %s\n", tag.OriginalFilename())
	fmt.Printf("  Playlist delay: %s\n", tag.PlaylistDelay())
	fmt.Printf("  Encoding time: %s\n", tag.BPM())
	fmt.Printf("  Original release time: %s\n", tag.OriginalReleaseTime())
	fmt.Printf("  Recording time: %s\n", tag.RecordingTime())
	fmt.Printf("  Release time: %s\n", tag.ReleaseTime())
	fmt.Printf("  Tagging time: %s\n", tag.TaggingTime())
	fmt.Printf("  Software/Hardware and settings used for encoding: %s\n", tag.EncodingSettings())
	fmt.Printf("  Album sort order: %s\n", tag.AlbumSortOrder())
	fmt.Printf("  Performer sort order: %s\n", tag.PerformerSortOrder())
	fmt.Printf("  Title sort order: %s\n", tag.TitleSortOrder())
	fmt.Println()
	fmt.Printf("  %d user defined text information frames\n", len(tag.UserTextFrames))
	for _, f := range tag.UserTextFrames {
		fmt.Printf("    %s: %s\n", f.Description, f.Value)
	}
}
