package main

import (
	"os"
	"strings"

	"testing"
)

const (
	username = "admin"
	password = "admin"
	ip       = "192.168.31.21"
	port     = 22
	cmd      = "show clock;exit"
)

func Test_SSH(t *testing.T) {
	session, err := connect(username, password, ip, port)
	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()

	cmdlist := strings.Split(cmd, ";")

	stdinBuf, err := session.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}
	//	var bt bytes.Buffer
	//	session.Stdout = &bt
	t.Log(session.Stdout)
	t.Log(session.Stderr)
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	err = session.Shell()
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cmdlist {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
	}
	session.Wait()
	t.Error(err)
	//	t.Log(bt.String())
	return
}
