package config

import (
	"os"
	"path/filepath"
)

func GetExecName() string {
	execName, err := os.Executable()
	if err != nil {
		return "unknown"
	}
	execName = filepath.Base(execName)
	return execName
}

func GetExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	return execPath
}
