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

	// registry map[string][]string
}

func (e *hashes) save(f file) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	defer e.wg.Done()

	h, err := hash(f.path)
	if err != nil {
		log.Fatal(err)
	}

	dups, ok := duplicates[h]
	var first string
	if ok {
		ndups = ndups + 1
		if len(dups) == 1 {
			ndups = ndups + 1
			first = dups[0]
		}
		if !quiet && showDups {
			if first != "" {
				printFirstDup(h, first)
			}
			printDup(h, f)
		}
	}
	dups = append(dups, f.id)
	duplicates[h] = dups
	// e.registry[h] = append(e.registry[h], f.id)
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
