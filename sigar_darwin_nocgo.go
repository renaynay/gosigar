// +build darwin
// +build !cgo

package gosigar

import (
	"runtime"
)

func (self *LoadAverage) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *Swap) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *HugeTLBPages) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *Uptime) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *Mem) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *Cpu) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *CpuList) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *FDUsage) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *FileSystemList) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcList) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcState) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcMem) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcTime) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcArgs) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcEnv) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}

func (self *ProcExe) Get(pid int) error {
	return ErrNotImplemented{runtime.GOOS}
}
