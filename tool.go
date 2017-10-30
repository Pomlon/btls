package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			fmt.Println("Executing default")
			buf := make([]byte, 2000)
			stdout.Read(buf)
			fmt.Println(string(buf))
		}
	}
}

func main() {
	proc := Proc{}
	fmt.Println("Executing")
	proc.ExecCommand("./courierApp.exe")

	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		fmt.Println("ForRun")

		if reader.Text() == "kill" {
			fmt.Println("Sending to chan")
			proc.ChanIn <- "kill"
		}

		if reader.Text() == "out" {
			os.Exit(0)
		}
	}
}
