package files

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// Copy source to destination path:
// if destination is a directory, source will be copied in the dir with the file basename.
func Copy(source, destination string) error {
	src := cleanPath(source)
	dst := cleanPath(destination)
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()

	if IsDir(src) {
		return fmt.Errorf("source path is a directory")
	}
	if IsDir(dst) {
		basename := filepath.Base(src)
		dst = path.Join(dst, basename)
	}

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

func existsWithError(filepath string) (bool, error) {
	name := cleanPath(filepath)
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
		// Windows: error 123 (0x7B) The filename, directory name, or volume label syntax is incorrect.
		if runtime.GOOS == "windows" {
			if e, ok := err.(*os.PathError); ok {
				if en, ok := e.Err.(syscall.Errno); ok {
					if int(en) == 123 {
						return false, err
					}
				}
			}
		}
		return true, err
	}
	return true, nil
}

// Exists reports whether the named file or directory exists.
func Exists(filepath string) bool {
	exist, _ := existsWithError(filepath)
	return exist
}

// IsAccessible reports if file exists and user has permission to use it
func IsAccessible(filepath string) bool {
	exist, err := existsWithError(filepath)
	return exist && !os.IsPermission(err)
}

// IsDir reports whether d is a directory.
func IsDir(fpath string) bool {
	d := cleanPath(fpath)
	if fi, err := os.Stat(d); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsRegular reports whether filepath is a regular file.
func IsRegular(fpath string) bool {
	f := cleanPath(fpath)
	if fi, err := os.Stat(f); err == nil {
		return fi.Mode().IsRegular()
	}
	return false
}

// Sha1Sum gives the checksum for the given file
func Sha1Sum(fpath string) (string, error) {
	name := cleanPath(fpath)
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	sha1 := sha1.New()
	_, err = io.Copy(sha1, reader)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha1.Sum(nil)), nil
}

// ReadLines returns a slice containing file lines.
func ReadLines(path string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return lines, err
}

/*
func MD5OfFile(fullpath string) []byte {
	if contents, err := ioutil.ReadFile(fullpath); err == nil {
		var md5sum hash.Hash = md5.New()
		md5sum.Write(contents)
		return md5sum.Sum()
	}
	return nil
}
*/

func cleanPath(filepath string) string {
	return strings.TrimSpace(filepath)
	/*
		withoutSpaces := strings.TrimSpace(filepath)
		if withoutSpaces == "" {
			return ""
		}
		return path.Clean(withoutSpaces)
	*/
}

// EachLineFunc is definition of callback
type EachLineFunc func(line string) error

// EachLine walks lines, calling EachLineFunc for each line of the file.
// All errors that arise visiting lines are filtered by callback function.
func EachLine(path string, walkFn EachLineFunc) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err = walkFn(line)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return err
}
