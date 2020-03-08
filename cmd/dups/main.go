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
	"github.com/enr/go-files/files"
)

var (
	processed  = 0
	ndups      = 0
	duplicates = make(map[string][]string)
	startTime  time.Time
	messages   = make(chan string)
	logger     = log.New(os.Stdout, "", 0)

	versionTemplate = `dups %s
Revision: %s
Build date: %s
`
	showDups        = true
	appVersion      string
	baseDirectory   string
	version         bool
	help            bool
	quiet           bool
	names           bool
	fullPath        bool
	ndupsAsExitCode bool
	excludes        []string

	// global wait group counting the probably duplicate files
	wg sync.WaitGroup
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
		logger.Printf("\nShutting down the program. At the moment: checked %d files and found %d dups\n", processed, ndups)
		os.Exit(0)
	}()

	defer func() {
		trace()
		exitCode := 0
		// exit code is the number of duplicates if quiet=true or
		// `ndupsAsExitCode` is explicitally set to true
		if ndupsAsExitCode || quiet {
			exitCode = ndups
		}
		os.Exit(exitCode)
	}()

	h := &hashes{
		mutex: new(sync.Mutex),
		wg:    new(sync.WaitGroup),
	}

	go func(*hashes) {
		processProbableDuplicate(h)
	}(h)

	flag.BoolVar(&version, "version", false, "show version")
	flag.BoolVar(&quiet, "quiet", false, "no output, exit code the number of dups")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&names, "names-only", false, "show only file names")
	flag.BoolVar(&fullPath, "full-path", false, "show full path for files")
	flag.BoolVar(&ndupsAsExitCode, "dups-exit", false, "set exit code to the number of duplicates")
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
	_, err := readDirectory(args[0])
	if err != nil {
		fmt.Printf("error %v\n", err)
		os.Exit(1)
	}
	wg.Wait()
	h.wg.Wait()
}

func readDirectory(dirpath string) (map[int64][]string, error) {
	return readDirectoryExcluding(dirpath, excludes)
}

func readDirectoryExcluding(dirpath string, excludes []string) (map[int64][]string, error) {
	sizemap := make(map[int64][]string)
	source, err := normalizePath(dirpath)
	if err != nil {
		return sizemap, err
	}
	if files.IsSymlink(source) {
		source, err = os.Readlink(source)
		if err != nil {
			return sizemap, err
		}
	}
	if !files.IsDir(source) {
		return sizemap, fmt.Errorf("%s not a directory", dirpath)
	}
	if !quiet && !names {
		logger.Printf("Looking for duplicates in %s", source)
	}
	baseDirectory = source
	err = filepath.Walk(source, func(fpath string, f os.FileInfo, err error) error {
		if !files.IsRegular(fpath) {
			return nil
		}
		s := f.Size()
		processed = processed + 1
		dups, ok := sizemap[s]
		if ok {
			wg.Add(1)
			messages <- fpath
			if len(dups) == 1 {
				wg.Add(1)
				messages <- dups[0]
			}
		}
		dups = append(dups, fpath)
		sizemap[s] = dups
		return nil
	})
	close(messages)
	return sizemap, err
}
