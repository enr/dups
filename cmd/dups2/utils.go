package main

import (
	"fmt"
	"math"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/enr/go-commons/lang"
)

func trace() {
	if quiet || names {
		return
	}
	endTime := time.Now()
	logger.Printf("Checked %d files and found %d dups in %s", processed, ndups, humanizeDuration(endTime.Sub(startTime)))
}

func normalizePath(dirpath string) (string, error) {
	p, err := filepath.Abs(dirpath)
	if err != nil {
		return "", err
	}
	p = filepath.ToSlash(p)
	return p, nil
}

func printFirstDup(checksum string, fp string) {
	f := fp
	if fullPath {
		f, _ = normalizePath(path.Join(baseDirectory, fp))
	}
	p(checksum, f)
}

func printDup(checksum string, fil file) {
	f := fil.id
	if fullPath {
		f = fil.path
	}
	p(checksum, f)
}

func p(checksum string, p string) {
	if names {
		logger.Println(p)
	} else {
		logger.Printf("%s %s", checksum, p)
	}
	//}
}

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65
func humanizeDuration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))
	millis := int64(math.Mod(float64(duration.Milliseconds()), 1000))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
		{"millisecond", millis},
	}

	parts := []string{}
	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	s := strings.TrimSpace(strings.Join(parts, " "))
	if s == "" {
		return "no time"
	}
	return s
}

func processProbableDuplicate(h *hashes) {

	for {
		fpath, ok := <-messages
		if ok == false {
			break
		}
		fileID, err := filepath.Rel(baseDirectory, fpath)
		if err != nil {
			break
		}
		fileID = filepath.ToSlash(fileID)
		if lang.SliceContainsString(excludes, fileID) {
			break
		}
		fullPath, err := normalizePath(fpath)
		if err != nil {
			break
		}
		fil := file{
			id:   fileID,
			path: fullPath,
		}
		h.wg.Add(1)
		h.save(fil)
	}
}
