package main

import (
	"errors"
	"net"
	"os"
	"time"
)

// stdioConn implements the Conn interface using standard input/output.
type stdioConn struct {
}

// NewStdioConn creates a new stdioConn.
func NewStdioConn() *stdioConn {
	return &stdioConn{}
}

func (s *stdioConn) Read(b []byte) (n int, err error) {
	return os.Stdin.Read(b)
}

func (s *stdioConn) Write(b []byte) (n int, err error) {
	return os.Stdout.Write(b)
}

func (s *stdioConn) Close() error {
	os.Exit(0)
	return nil
}

func (s *stdioConn) LocalAddr() net.Addr {
	return nil // Local address not applicable for stdio
}

func (s *stdioConn) RemoteAddr() net.Addr {
	return nil // Remote address not applicable for stdio
}

func (s *stdioConn) SetDeadline(t time.Time) error {
	return errors.New("deadline not supported for stdioConn")
}

func (s *stdioConn) SetReadDeadline(t time.Time) error {
	return errors.New("deadline not supported for stdioConn")
}

func (s *stdioConn) SetWriteDeadline(t time.Time) error {
	return errors.New("deadline not supported for stdioConn")
}
