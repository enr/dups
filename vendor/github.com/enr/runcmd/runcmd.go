package runcmd

import (
	"bytes"
	_ "fmt"
	"os/exec"
	"syscall"
)

// Streams provides access to the output and error buffers.
type Streams struct {
	out *bytes.Buffer
	err *bytes.Buffer
}

// Stderr returns the underlying buffer with the contents of the error stream.
func (s *Streams) Stderr() *bytes.Buffer {
	return s.err
}

// Stdout returns the underlying buffer with the contents of the output stream.
func (s *Streams) Stdout() *bytes.Buffer {
	return s.out
}

func getExitStatus(err error) int {
	if err == nil {
		return 0
	}
	exitStatus := 1
	if msg, ok := err.(*exec.ExitError); ok {
		exitStatus = msg.Sys().(syscall.WaitStatus).ExitStatus()
	}
	return exitStatus

	// if ee, ok := err.(*exec.ExitError); ok && err != nil {
	// 	status := ee.ProcessState.Sys().(syscall.WaitStatus)
	// 	if status.Exited() {
	// 		// A non-zero return code isn't considered an error here.
	// 		result.Code = status.ExitStatus()
	// 		err = nil
	// 	}
	// 	logger.Infof("run result: %v", ee)
	// }
}
