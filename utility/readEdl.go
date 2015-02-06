package utility

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type EdlShot struct {
	ShotCode      string
	ShotType      string
	StartDateTime string
	EndDateTime   string
	FromClipName  string
	SourceFile    string
}

func ReadEdl(filePth string) ([]*EdlShot, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	var shots []*EdlShot
	var s *EdlShot
	flag := 0
	for {
		line, err := bfRd.ReadBytes('\n')
		if v, _ := regexp.MatchString("^[0-9]+$", string(line[:3])); v {
			if flag > 0 {
				shots = append(shots, s)
			}
			s = &EdlShot{}
			s.ShotCode = string(line[:3])
			s.StartDateTime = string(line[29:40])
			s.EndDateTime = string(line[41:52])

			flag += 1
		} else if string(line[2:6]) == "FROM" {
			s.FromClipName = strings.TrimSpace(string(line[19:]))
		} else if string(line[2:8]) == "SOURCE" {
			s.SourceFile = strings.TrimSpace(string(line[15:]))
		} else if string(line[0:2]) == "* " {
			s.ShotType = strings.TrimSpace(string(line[2:]))
		}

		if err != nil {
			if err == io.EOF {
				shots = append(shots, s)
			}
			return shots, err
		}
	}
	return shots, nil
}
