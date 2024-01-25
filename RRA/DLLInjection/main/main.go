package main

import (
	"syscall"

	"github.com/JamesHovious/w32"
)

func main() {
	dllPath, _ := syscall.FullPath("DTest.dll")

	className, _ := syscall.UTF16PtrFromString("Notepad")
	hwnd := w32.FindWindowW(className, nil)
	_, processId := w32.GetWindowThreadProcessId(hwnd)

	var dwMemSize int
	var hProc w32.HANDLE
	var err error
	var lpRemoteRem, lpLoadLibrary uintptr

	hProc, err = w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, uint32(processId))
	if err == nil {
		dwMemSize = len(dllPath) + 1
		lpRemoteRem, err = w32.VirtualAllocEx(hProc, 0, dwMemSize, w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_READWRITE)
		if err == nil {
			err = w32.WriteProcessMemory(hProc, uint32(lpRemoteRem), []byte(dllPath), uint(dwMemSize))
			if err == nil {
				modulKernel, _ := syscall.LoadLibrary("kernel32.dll")
				lpLoadLibrary, err = syscall.GetProcAddress(modulKernel, "LoadLibraryA")
				if err == nil {
					hTread, _, err := w32.CreateRemoteThread(hProc, nil, 0, uint32(lpLoadLibrary), lpRemoteRem, 0)
					if err == nil {
						w32.ResumeThread(hTread)
						w32.WaitForSingleObject(hTread, syscall.INFINITE)
						_, err := w32.GetExitCodeProcess(hProc)
						if err == nil {
							w32.CloseHandle(hTread)
						} else {
							panic(err)
						}
					} else {

					}
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			w32.VirtualFreeEx(hProc, lpRemoteRem, 0, w32.MEM_RELEASE)
			panic(err)
		}
	} else {
		panic(err)
	}

	w32.CloseHandle(hProc)

}
