package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"

	".."
)

func main() {
	exitcode := 0
	for _, path := range os.Args[1:] {
		err := rename(path)
		if err != nil {
			log.Printf("%s: %s", path, err)
			exitcode++
		}
	}
	os.Exit(exitcode)
}

func rename(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tag, err := id3v2.Read(file)
	if err != nil {
		return err
	}

	fname, err := format(tag)
	if err != nil {
		return err
	}
	dest := filepath.Join(filepath.Dir(path), fname)
	fmt.Printf("%s -> %s\n", path, dest)
	return os.Rename(path, dest)
}

func format(tag *id3v2.Tag) (string, error) {
	artist := tag.Artist()

	title := tag.Title()
	if title == "" {
		return "", fmt.Errorf("Empty title")
	}

	track := tag.Track()
	if idx := strings.Index(track, "/"); idx != -1 {
		track = track[0:idx]
	}
	if n, err := strconv.Atoi(track); err == nil {
		track = fmt.Sprintf("%02d", n)
	}
	base := ""
	if artist == "" {
		base = fmt.Sprintf("%s %s.mp3", track, title)
	} else {
		base = fmt.Sprintf("%s [%s] %s.mp3", track, artist, title)
	}
	return safeFilename(toHankaku(base)), nil
}

func toHankaku(s string) string {
	buf := make([]rune, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		buf[i] = hankaku(r)
		i++
	}
	return string(buf)
}

func hankaku(r rune) rune {
	if r == '　' {
		return ' '
	} else if 'Ａ' <= r && r <= 'ｚ' {
		return r - 'Ａ' + 'A'
	} else {
		return r
	}
}

func safeFilename(s string) string {
	return strings.Replace(s, "/", "／", -1)

}
