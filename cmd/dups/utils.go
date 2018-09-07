package main

import (
	"path/filepath"
	"time"
)

func trace(s string) (string, time.Time) {
	//logger.Printf("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	if quiet || names {
		return
	}
	endTime := time.Now()
	//logger.Println("  END:", s, "ElapsedTime:", )
	logger.Printf("Checked %d files and found %d dups in %s", processed, ndups, endTime.Sub(startTime))
}

func normalizePath(dirpath string) (string, error) {
	p, err := filepath.Abs(dirpath)
	if err != nil {
		return "", err
	}
	p = filepath.ToSlash(p)
	return p, nil
}

func printDups(checksum string, dups []string) {
	if !showDups {
		return
	}
	for _, f := range dups {
		if fullPath {
			f, _ = normalizePath(f)
		}
		if names {
			logger.Println(f)
		} else {
			logger.Printf("%s %s", checksum, f)
		}
	}
}
