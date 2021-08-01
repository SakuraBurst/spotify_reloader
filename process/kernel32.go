package process

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	create32Snap    = kernel32.NewProc("CreateToolhelp32Snapshot")
	openProceess    = kernel32.NewProc("OpenProcess")
	processFirst    = kernel32.NewProc("Process32First")
	processNext     = kernel32.NewProc("Process32Next")
	moduleFirst     = kernel32.NewProc("Module32First")
	moduleNext      = kernel32.NewProc("Module32Next")
	procCloseHandle = kernel32.NewProc("CloseHandle")
	killProcess     = kernel32.NewProc("TerminateProcess")
)

func CreateToolhelp32Snapshot(flag, id uint32) uintptr {
	snap, _, _ := create32Snap.Call(uintptr(flag), uintptr(id))
	return snap
}

func Process32First(snap uintptr, processEntry *PROCESSENTRY32) bool {
	ret, _, _ := processFirst.Call(snap, uintptr(unsafe.Pointer(processEntry)))
	fmt.Println(ret)
	return ret != 0
}

func Process32Next(snap uintptr, processEntry *PROCESSENTRY32) bool {
	ret, _, _ := processNext.Call(snap, uintptr(unsafe.Pointer(processEntry)))
	return ret != 0
}

func Module32First(snap uintptr, moduleEntry *MODULEENTRY32) bool {
	ret, _, _ := moduleFirst.Call(snap, uintptr(unsafe.Pointer(moduleEntry)))
	return ret != 0
}

func Module32Next(snap uintptr, moduleEntry *MODULEENTRY32) bool {
	ret, _, _ := moduleNext.Call(snap, uintptr(unsafe.Pointer(moduleEntry)))
	return ret != 0
}

func OpenProcess(flag, id uint32) (HANDLE, bool) {
	ret, _, _ := openProceess.Call(uintptr(flag), 0, uintptr(id))

	return HANDLE(ret), ret != 0
}

func CloseHandle(handle HANDLE) {
	procCloseHandle.Call(uintptr(handle))
}

func KillProcess(handle HANDLE, exitCode uint32) bool {
	ret, _, _ := killProcess.Call(uintptr(handle), uintptr(exitCode))
	return ret != 0
}
