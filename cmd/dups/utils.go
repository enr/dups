package main

import (
	"fmt"
	"math"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func trace() {
	if quiet || names {
		return
	}
	endTime := time.Now()
	//logger.Println("  END:", s, "ElapsedTime:", )
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

func printDups2(checksum string, fp string) {
	if !showDups {
		return
	}
	//for _, f := range dups {
	f := fp
	if fullPath {
		f, _ = normalizePath(path.Join(baseDirectory, fp))
	}
	if names {
		logger.Println(f)
	} else {
		logger.Printf("%s %s", checksum, f)
	}
	//}
}

func printDups(checksum string, fil file) {
	if !showDups {
		return
	}
	//for _, f := range dups {
	f := fil.id
	if fullPath {
		f = fil.path
	}
	if names {
		logger.Println(f)
	} else {
		logger.Printf("%s %s", checksum, f)
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

	return strings.Join(parts, " ")
}
