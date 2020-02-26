package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type hashes struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup

	registry map[string][]string
}

func (e *hashes) save(f file) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	defer e.wg.Done()

	processed = processed + 1
	h, err := hash(f.path)
	if err != nil {
		log.Fatal(err)
	}

	dups, ok := e.registry[h]
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
				printFirstDup(h, first)
			}
			printDups(h, f)
		}
	}
	duplicates[h] = dups

	e.registry[h] = append(e.registry[h], f.id)
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
