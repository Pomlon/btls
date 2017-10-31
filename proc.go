package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Proc struct {
	ChanIn  chan string
	Running bool
}

func (p *Proc) ExecCommand(command string) {
	p.ChanIn = make(chan string, 100)
	go p.Command(command, p.ChanIn)
}

func (p *Proc) Command(command string, cha <-chan string) {
	cmd := exec.Command(command)

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdout.Close()

	time.Sleep(time.Millisecond * 1000)
	p.Running = true
	for p.Running == true {
		//fmt.Println("SELECTIN")
		select {
		case msgIn := <-cha:
			if msgIn == "kill" {
				stdout.Close()
				cmd.Process.Kill()
				p.Running = false
			}

		default:

			//fmt.Println("Executing default")
			buf := make([]byte, 2000)
			n, _ := stdout.Read(buf)
			if str := string(buf); n > 0 {
				fmt.Println(str)
			}

		}
	}
}

//Murder murders the running process
func (p *Proc) Murder() {
	p.ChanIn <- "kill"
}

//OneRun runs command once, waits for it and returns output.
func OneRun(command string) {
	splat := strings.Fields(command)
	comms := splat[1:]
	out, err := exec.Command(splat[0], comms...).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}

//OneRunMany same as OneRun but accepts many commands
func OneRunMany(commands ...string) {
	for _, comm := range commands {
		OneRun(comm)
	}
}
