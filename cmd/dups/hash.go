package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// Executor ...
type Executor struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup

	Results map[string][]string
}

// SaveFileHash Collect all files and save their hashes to the mapping
func (e *Executor) SaveFileHash(f file) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	defer e.wg.Done()

	processed = processed + 1
	h, err := hash(f.path)
	if err != nil {
		log.Fatal(err)
	}

	dups, ok := e.Results[h]
	var first string
	if len(dups) == 1 {
		first = dups[0]
	}
	dups = append(dups, f.id)
	if ok {
		ndups = ndups + 1
		if !quiet {
			if first != "" {
				ndups = ndups + 1
				printDups2(h, first)
			}
			printDups(h, f)
		}
	}
	duplicates[h] = dups

	e.Results[h] = append(e.Results[h], f.id)
}

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
