package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
	//	"github.com/bitly/go-simplejson"
)

type sshhost struct {
	host     string
	port     int
	username string
	password string
	cmd      []string
}

func main() {
	hosts := flag.String("hosts", "", "host address list")
	cmd := flag.String("cmd", "", "cmds")
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	port := flag.Int("port", 22, "ssh port")
	cmdfile := flag.String("cmdfile", "", "cmdfile path")
	hostfile := flag.String("hostfile", "", "hostfile path")
	ipfile := flag.String("ipfile", "", "hostfile path")
	cfg := flag.String("cfg", "", "cfg path")

	flag.Parse()

	var cmdlist []string
	var hostlist []string
	var err error

	sshhosts := []sshhost{}
	var host_struct sshhost

	if *ipfile != "" {
		hostlist, err = GetIpList(*ipfile)
		if err != nil {
			log.Println("load hostlist error: ", err)
			return
		}
	}

	if *hostfile != "" {
		hostlist, err = Getfile(*hostfile)
		if err != nil {
			log.Println("load hostfile error: ", err)
			return
		}
	}
	if *hosts != "" {
		hostlist = strings.Split(*hosts, ";")

	}

	if *cmdfile != "" {
		cmdlist, err = Getfile(*cmdfile)
		if err != nil {
			log.Println("load cmdfile error: ", err)
			return
		}
	}
	if *cmd != "" {
		cmdlist = strings.Split(*cmd, ";")
	}

	if *cfg == "" {
		for _, host := range hostlist {
			host_struct.host = host
			host_struct.username = *username
			host_struct.password = *password
			host_struct.port = *port
			host_struct.cmd = cmdlist
			sshhosts = append(sshhosts, host_struct)
		}
	}
	/*
		 else {
			cfgjson, err := GetfileAll(*cfg)
			if err != nil {
				log.Println("load cfg error: ", err)
				return
			}

				js, js_err := simplejson.NewJson(cfgjson)
				if js_err != nil {
					log.Println("json format error: ", js_err)
					return
				}


		}
	*/
	//fmt.Println(sshhosts)

	chs := make([]chan string, len(sshhosts))
	for i, host := range sshhosts {
		chs[i] = make(chan string, 1)
		go dossh(host.username, host.password, host.host, host.cmd, host.port, chs[i])
	}
	for i, ch := range chs {
		fmt.Println(sshhosts[i].host, " ssh start")
		select {
		case res := <-ch:
			if res != "" {
				fmt.Println(res)
			}
		case <-time.After(30 * 1000 * 1000 * 1000):
			log.Println("SSH run timeout")
		}
		fmt.Println(sshhosts[i].host, " ssh end\n")
	}

}
