package process

type HANDLE uintptr

type PROCESSENTRY32 struct {
	dwSize              uint32
	cntUsage            uint32
	th32ProcessID       uint32
	th32DefaultHeapID   uintptr
	th32ModuleID        uint32
	cntThreads          uint32
	th32ParentProcessID uint32
	pcPriClassBase      uint64
	dwFlags             uint32
	szExeFile           [MAX_PATH]byte
}

type MODULEENTRY32 struct {
	dwSize        uint32
	th32ModuleID  uint32
	th32ProcessID uint32
	glblcntUsage  uint32
	proccntUsage  uint32
	modBaseAddr   *byte
	modBaseSize   uint32
	hModule       HANDLE
	szModule      [MAX_MODULE_NAME32 + 1]byte
	szExeFile     [MAX_PATH]byte
}
