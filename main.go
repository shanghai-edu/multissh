package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/shanghai-edu/multissh/funcs"
	"github.com/shanghai-edu/multissh/g"
)

func main() {
	version := flag.Bool("v", false, "show version")
	hosts := flag.String("hosts", "", "host address list")
	ips := flag.String("ips", "", "ip address list")
	cmds := flag.String("cmds", "", "cmds")
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	key := flag.String("k", "", "ssh private key")
	port := flag.Int("port", 22, "ssh port")
	ciphers := flag.String("ciphers", "", "ciphers")
	cmdFile := flag.String("cmdfile", "", "cmdfile path")
	hostFile := flag.String("hostfile", "", "hostfile path")
	ipFile := flag.String("ipfile", "", "ipfile path")
	cfgFile := flag.String("c", "", "cfg File Path")
	jsonMode := flag.Bool("j", false, "print output in json format")
	outTxt := flag.Bool("outTxt", false, "write result into txt")
	fileLocate := flag.String("f", "", "write file locate")
	linuxMode := flag.Bool("l", false, "In linux mode,multi command combine with && ,such as date&&cd /opt&&ls")
	timeLimit := flag.Int("t", 30, "max timeout")
	numLimit := flag.Int("n", 20, "max execute number")

	flag.Parse()

	var cmdList, hostList, cipherList []string
	var err error

	sshHosts := []g.SSHHost{}
	var host_Struct g.SSHHost

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *ipFile != "" {
		hostList, err = g.GetIpListFromFile(*ipFile)
		if err != nil {
			log.Println("load iplist error: ", err)
			return
		}
	}

	if *hostFile != "" {
		hostList, err = g.Getfile(*hostFile)
		if err != nil {
			log.Println("load hostfile error: ", err)
			return
		}
	}
	if *ips != "" {
		hostList, err = g.GetIpList(*ips)
		if err != nil {
			log.Println("load iplist error: ", err)
			return
		}
	}

	if *hosts != "" {
		hostList = g.SplitString(*hosts)
	}

	if *cmdFile != "" {
		cmdList, err = g.Getfile(*cmdFile)
		if err != nil {
			log.Println("load cmdfile error: ", err)
			return
		}
	}
	if *cmds != "" {
		cmdList = g.SplitString(*cmds)

	}
	if *ciphers != "" {
		cipherList = g.SplitString(*ciphers)
	}
	if *cfgFile == "" {
		for _, host := range hostList {
			host_Struct.Host = host
			host_Struct.Username = *username
			host_Struct.Password = *password
			host_Struct.Port = *port
			host_Struct.CmdList = cmdList
			host_Struct.Key = *key
			host_Struct.LinuxMode = *linuxMode
			sshHosts = append(sshHosts, host_Struct)
		}
	} else {
		sshHosts, err = g.GetJsonFile(*cfgFile)
		if err != nil {
			log.Println("load cfgFile error: ", err)
			return
		}
		for i := 0; i < len(sshHosts); i++ {
			if sshHosts[i].Cmds != "" {
				sshHosts[i].CmdList = g.SplitString(sshHosts[i].Cmds)
			} else {
				cmdList, err = g.Getfile(sshHosts[i].CmdFile)
				if err != nil {
					log.Println("load cmdFile error: ", err)
					return
				}
				sshHosts[i].CmdList = cmdList
			}
		}
	}

	chLimit := make(chan bool, *numLimit) //控制并发访问量
	chs := make([]chan g.SSHResult, len(sshHosts))
	startTime := time.Now()
	log.Println("Multissh start")
	limitFunc := func(chLimit chan bool, ch chan g.SSHResult, host g.SSHHost) {
		funcs.Dossh(host.Username, host.Password, host.Host, host.Key, host.CmdList, host.Port, *timeLimit, cipherList, host.LinuxMode, ch)
		<-chLimit
	}
	for i, host := range sshHosts {
		chs[i] = make(chan g.SSHResult, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], host)
	}
	sshResults := []g.SSHResult{}
	for _, ch := range chs {
		res := <-ch
		if res.Result != "" {
			sshResults = append(sshResults, res)
		}
	}
	endTime := time.Now()
	log.Printf("Multissh finished. Process time %s. Number of active ip is %d", endTime.Sub(startTime), len(sshHosts))
	//gu
	if *outTxt {
		for _, sshResult := range sshResults {
			err = g.WriteIntoTxt(sshResult, *fileLocate)
			if err != nil {
				log.Println("write into txt error: ", err)
				return
			}
		}
		return
	}
	if *jsonMode {
		jsonResult, err := json.Marshal(sshResults)
		if err != nil {
			log.Println("json Marshal error: ", err)
		}
		fmt.Println(string(jsonResult))
		return
	}
	for _, sshResult := range sshResults {
		fmt.Println("host: ", sshResult.Host)
		fmt.Println("========= Result =========")
		fmt.Println(sshResult.Result)
	}

}
