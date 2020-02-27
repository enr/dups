package qac

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// FullExePathOrFail returns the absolute and normalized path to the exe with conventional extensione (ie ".exe" for Windows, "" for other systems) and calls log.Fatal if it returns a non-nil error.
func FullExePathOrFail(p string) string {
	fp, err := FullExePath(p)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

// FullExePath returns:
// - the absolute and normalized path to the exe with conventional extensione (ie ".exe" for Windows, "" for other systems)
// - error if path does not exists
func FullExePath(p string) (string, error) {
	// adjusted := fmt.Sprintf("../../bin/%s", p)
	abs, _ := filepath.Abs(p)
	fp := filepath.FromSlash(filepath.Clean(abs))
	var ext string
	if runtime.GOOS == "windows" && isMissingExt(fp) {
		ext = ".exe"
	}
	executablePath := fmt.Sprintf("%s%s", fp, ext)
	if _, err := os.Stat(executablePath); os.IsNotExist(err) {
		return "", fmt.Errorf("no such executable: %s", executablePath)
	}
	return executablePath, nil
}

func isMissingExt(p string) bool {
	return !strings.HasSuffix(p, ".exe") && !strings.HasSuffix(p, ".bat") && !strings.HasSuffix(p, ".cmd")
}
