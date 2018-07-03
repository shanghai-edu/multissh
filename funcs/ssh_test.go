package funcs

import (
	"bytes"
	//	"os"
	"strings"
	"testing"
)

const (
	username = "root"
	password = ""
	ip       = "192.168.80.131"
	port     = 22
	cmd      = "cd /opt;pwd;exit"
	key      = "../server.key"
)

func Test_SSH(t *testing.T) {
	var cipherList []string
	session, err := connect(username, password, ip, key, port, cipherList)
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

	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt

	session.Stderr = &errbt
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
	t.Log((outbt.String() + errbt.String()))
	return
}

/*
func Test_SSH_run(t *testing.T) {
	var cipherList []string
	session, err := connect(username, password, ip, key, port, cipherList)
	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()

	//cmdlist := strings.Split(cmd, ";")
	//newcmd := strings.Join(cmdlist, "&&")
	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt

	session.Stderr = &errbt
	err = session.Run(cmd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log((outbt.String() + errbt.String()))

	return
}
*/
