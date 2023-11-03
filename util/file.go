package util

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ExeDir() string {
	execFile, _ := exec.LookPath(os.Args[0])
	absPath, _ := filepath.Abs(execFile)
	exeDir := filepath.Dir(absPath)
	return exeDir
}
