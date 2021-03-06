// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"github.com/ysicing/go-utils/exfile"
	"runtime"
)

// IsMacOS is Mac OS
func IsMacOS() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsLinux is linux
func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

// IsUnix macos or linux
func IsUnix() bool {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsContainer 是否是容器
func IsContainer() bool {
	return exfile.CheckFileExistsv2("/.dockerenv")
}
