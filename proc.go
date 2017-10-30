package main

import (
	"fmt"
	"os/exec"
	"strings"
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

	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()
	cmd.Start()
	p.Running = true
	for p.Running == true {
		fmt.Println("SELECTIN")
		select {
		case msgIn := <-cha:
			if msgIn == "kill" {
				cmd.Process.Kill()
				p.Running = false
			}

		default:
			if cmd.ProcessState.Exited() == false {
				fmt.Println("Executing default")
				buf := make([]byte, 2000)
				stdout.Read(buf)
				fmt.Println(string(buf))
			} else {
				break
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
