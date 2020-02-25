package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/enr/dups/lib/core"
	"github.com/enr/go-commons/lang"
	"github.com/enr/go-files/files"
)

var (
	processed  = 0
	ndups      = 0
	duplicates = make(map[string][]string)
	logger     = log.New(os.Stdout, "", 0)
	showDups   = true

	versionTemplate = `dups %s
Revision: %s
Build date: %s
`
	appVersion    string
	baseDirectory string
	version       bool
	help          bool
	quiet         bool
	names         bool
	fullPath      bool

	startTime time.Time
)

type file struct {
	id   string
	path string
}

func main() {

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGKILL, syscall.SIGHUP)
		defer close(signalChan)

		<-signalChan
		logger.Println("Shutting down the program...")
		os.Exit(0)
	}()

	defer func() {
		trace()
		os.Exit(ndups)
	}()

	flag.BoolVar(&version, "version", false, "show version")
	flag.BoolVar(&quiet, "quiet", false, "no output, exit code the number of dups")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&names, "names-only", false, "show only file names")
	flag.BoolVar(&fullPath, "full-path", false, "show full path for files")
	flag.Parse()

	appVersion = fmt.Sprintf(versionTemplate, core.Version, core.GitCommit, core.BuildTime)

	if help {
		fmt.Printf(appVersion)
		flag.PrintDefaults()
		os.Exit(0)
	}
	if version {
		fmt.Printf(appVersion)
		os.Exit(0)
	}
	startTime = time.Now()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("error missing path")
		os.Exit(1)
	}
	baseDirectory = args[0]
	_, err := readDirectory(baseDirectory)
	if err != nil {
		fmt.Printf("error %v\n", err)
		os.Exit(1)
	}
}

func readDirectory(dirpath string) (map[string][]string, error) {
	return readDirectoryExcluding(dirpath, []string{})
}

func readDirectoryExcluding(dirpath string, excludes []string) (map[string][]string, error) {

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

	exec := &Executor{
		mutex:   new(sync.Mutex),
		wg:      new(sync.WaitGroup),
		Results: make(map[string][]string),
	}

	err = filepath.Walk(source, func(fpath string, f os.FileInfo, err error) error {
		if !files.IsRegular(fpath) {
			return nil
		}
		fileID, err := filepath.Rel(source, fpath)
		if err != nil {
			return err
		}
		fileID = filepath.ToSlash(fileID)
		if lang.SliceContainsString(excludes, fileID) {
			return nil
		}
		fullPath, err := normalizePath(fpath)
		// normalizePath(path.Join(baseDirectory, fpath))
		if err != nil {
			return err
		}

		fil := file{
			id:   fileID,
			path: fullPath,
		}

		exec.wg.Add(1)
		go exec.SaveFileHash(fil)

		// checksum, err := hash(fullPath)
		// if err != nil {
		// 	return err
		// }
		// processed = processed + 1
		// dups, ok := duplicates[checksum]
		// dups = append(dups, fileID)
		// if ok {
		// 	ndups = ndups + 1
		// 	if !quiet {
		// 		printDups(checksum, dups)
		// 	}
		// }
		// duplicates[checksum] = dups
		return nil
	})
	exec.wg.Wait()
	return duplicates, err
}
