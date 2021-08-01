package process

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"unsafe"
)

type Process struct {
	PID    uint32
	Handle HANDLE
	Name   string
	PPath  string
}

func (p *Process) OpenProcess() {
	proc, isNotErr := OpenProcess(PROCESS_ALL_ACCESS, p.PID)
	if !isNotErr {
		log.Fatal("process opening fail")
	}
	fmt.Println(proc)
	p.Handle = proc

}

func (p *Process) DeleteProcess() {
	isNotErr := KillProcess(p.Handle, 2)
	if !isNotErr {
		fmt.Println("ussuccesfuly kill proccess. id: ", p.PID)
	}
	fmt.Println("succesfuly kill proccess. id: ", p.PID)

}

func FindProcessesByName(name string) []Process {
	snap := CreateToolhelp32Snapshot(TH32CS_SNAPALL, 0)
	defer CloseHandle(HANDLE(snap))
	processes := make([]Process, 0)
	var pe32 PROCESSENTRY32
	pe32.dwSize = uint32(unsafe.Sizeof(pe32))
	fmt.Println(pe32.dwSize)
	if Process32First(snap, &pe32) {
		for Process32Next(snap, &pe32) {
			proc, err := GetProcessFullInfo(pe32.th32ProcessID)
			if err == nil && proc.Name == name {
				processes = append(processes, proc)
			}
		}
	} else {
		log.Fatal("хахаха приокл")
	}
	if len(processes) == 0 {
		log.Fatal(name, " is not founded")
	}
	for i := range processes {
		processes[i].OpenProcess()
	}
	fmt.Println(processes)
	return processes
}

func GetProcessFullInfo(pid uint32) (Process, error) {
	var me32 MODULEENTRY32
	me32.dwSize = uint32(unsafe.Sizeof(me32))
	snap := CreateToolhelp32Snapshot(TH32CS_MODULE, pid)
	defer CloseHandle(HANDLE(snap))
	if Module32First(snap, &me32) {
		return Process{PID: pid, Name: string(me32.szModule[:bytes.IndexByte(me32.szModule[:], 0)]), PPath: string(me32.szExeFile[:bytes.IndexByte(me32.szExeFile[:], 0)])}, nil
	} else {
		return Process{}, errors.New("cannot find module by pid")
	}

}
