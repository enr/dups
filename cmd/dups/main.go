package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/enr/dups/lib/colorgen"
	"github.com/enr/dups/lib/core"
	"github.com/enr/go-files/files"
	"github.com/spf13/pflag"
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
	excludes        GlobArray
	includes        GlobArray

	// global wait group counting the probably duplicate files
	wg sync.WaitGroup
)

type file struct {
	id   string
	path string
}

func init() {

	pflag.BoolVar(&version, "version", false, "show version")
	pflag.BoolVarP(&quiet, "quiet", "q", false, "no output, exit code the number of dups")
	pflag.BoolVarP(&help, "help", "h", false, "show help")
	pflag.BoolVarP(&names, "names-only", "n", false, "show only file names")
	pflag.BoolVarP(&fullPath, "full-path", "f", false, "show full path for files")
	pflag.BoolVarP(&ndupsAsExitCode, "dups-exit", "E", false, "set exit code to the number of duplicates")
	pflag.VarP(&excludes, "exclude", "e", "exclude filename glob patterns (can be supplied multiple times)")
	pflag.VarP(&includes, "include", "i", "only include filename glob patterns (can be supplied multiple times)")

	// These flags are automatically read by go-supportscolor
	// https://github.com/jwalton/go-supportscolor#info
	pflag.Bool("color", false, "force colored output")
	pflag.Bool("no-color", false, "disables colored output")

	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, appVersion)
		fmt.Fprintln(os.Stderr, "Flags:")
		pflag.PrintDefaults()
	}

	appVersion = fmt.Sprintf(versionTemplate, core.Version, core.GitCommit, core.BuildTime)
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

	rep := &reporter{
		color: colorgen.NewGenerator(),
	}
	h := &hashes{
		mutex: new(sync.Mutex),
		wg:    new(sync.WaitGroup),
		rep:   rep,
	}

	go func(*hashes) {
		processProbableDuplicate(h)
	}(h)

	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Print(appVersion)
		os.Exit(0)
	}
	startTime = time.Now()
	var startDir string
	args := pflag.Args()
	switch len(args) {
	case 0:
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get current directory: %s\n", err)
			os.Exit(1)
		}
		startDir = wd
	case 1:
		startDir = args[0]
	default:
		fmt.Fprintf(os.Stderr, "Error: only one path may be supplied\n")
		os.Exit(1)
	}
	_, err := readDirectory(startDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", startDir, err)
		os.Exit(1)
	}
	wg.Wait()
	h.wg.Wait()
}

func readDirectory(dirpath string) (map[int64][]string, error) {
	sizemap := make(map[int64][]string)
	source := filepath.Clean(dirpath)
	if files.IsSymlink(source) {
		var err error
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
	err := filepath.Walk(source, func(fpath string, f os.FileInfo, err error) error {
		if excludes.Match(f.Name()) {
			return filepath.SkipDir
		}
		if len(includes) > 0 {
			if !includes.Match(f.Name()) {
				// No SkipDir here, as we might want to include a child path
				return nil
			}
		}
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
