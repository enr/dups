package main

/*

Find duplicate files (same xxhash) in a directory.

Usage:

go run main.go ~/Pictures/

TODO:

* aggiungere option --exclude=.git/,/node_modules/
  dove fileID, err := filepath.Rel(source, fpath) aggiungere il ToSlash per confronto
  e dove lang.SliceContainsString(excludes, fileID) usare strings.Contains()
* riprovare benchmark in golang/dups_bench perche' almeno su linux meglio sha1 che l'altro hash usato

*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/enr/go-commons/lang"
	"github.com/enr/go-files/files"
)

var (
	processed = 0
	ndups     = 0
	logger    = log.New(os.Stdout, "", 0)
	showDups  = true

	versionTemplate = `dups %s
Revision: %s
Build date: %s
`
	appVersion string
	version    bool
	help       bool
	quiet      bool
	names      bool
	fullPath   bool
)

func main() {
	defer func() {
		un(trace("dups"))
		os.Exit(ndups)
	}()

	flag.BoolVar(&version, "version", false, "show version")
	flag.BoolVar(&quiet, "quiet", false, "no output, exit code the number of dups")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&names, "names-only", false, "show only file names")
	flag.BoolVar(&fullPath, "full-path", false, "show full path for files")
	flag.Parse()

	appVersion = fmt.Sprintf(versionTemplate, Version, GitCommit, BuildTime)

	if help {
		fmt.Printf(appVersion)
		flag.PrintDefaults()
		os.Exit(0)
	}
	if version {
		fmt.Printf(appVersion)
		os.Exit(0)
	}
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("error missing path")
		os.Exit(1)
	}
	dir := args[0]
	_, err := readDirectory(dir)
	if err != nil {
		fmt.Printf("error %v\n", err)
		os.Exit(1)
	}
	//printFinalReport(processed, ndups)

}

func readDirectory(dirpath string) (map[string][]string, error) {
	return readDirectoryExcluding(dirpath, []string{})
}

func readDirectoryExcluding(dirpath string, excludes []string) (map[string][]string, error) {
	duplicates := make(map[string][]string)
	source, err := normalizePath(dirpath)
	if err != nil {
		return duplicates, err
	}
	if files.IsSymlink(source) {
		source, err = os.Readlink(source)
		if err != nil {
			return duplicates, err
		}
	}
	if !files.IsDir(source) {
		return duplicates, fmt.Errorf("%s not a directory", dirpath)
	}
	if !quiet && !names {
		logger.Printf("Looking for duplicates in %s", source)
	}
	err = filepath.Walk(source, func(fpath string, f os.FileInfo, err error) error {
		if !files.IsRegular(fpath) {
			return nil
		}
		fileID, err := filepath.Rel(source, fpath)
		if err != nil {
			return err
		}
		if lang.SliceContainsString(excludes, fileID) {
			return nil
		}
		fullPath, err := normalizePath(fpath)
		if err != nil {
			return err
		}
		checksum, err := hash(fullPath)
		if err != nil {
			return err
		}
		processed = processed + 1
		dups, ok := duplicates[checksum]
		dups = append(dups, fileID)
		if ok {
			ndups = ndups + 1
			if !quiet {
				printDups(checksum, dups)
			}
		}
		duplicates[checksum] = dups
		return nil
	})
	return duplicates, err
}
