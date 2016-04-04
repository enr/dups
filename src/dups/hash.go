package main

/*

Find duplicate files (same xxhash) in a directory.

Usage:

go run main.go ~/Pictures/
*/

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

// Size of a SHA1 checksum in bytes.
const Size = 20

func hash(fullpath string) (string, error) {
	fh, err := os.Open(fullpath)
	defer fh.Close()
	if err != nil {
		return "", err
	}
	h := sha1.New()
	io.Copy(h, fh)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
