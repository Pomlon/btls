package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Proc struct {
	ChanIn chan string
	proc   *exec.Cmd
}

func (p *Proc) ExecCommand(command string) {
	p.proc = exec.Command(command)

	p.proc.Stdout = os.Stdout
	p.proc.Stderr = os.Stderr

	err := p.proc.Start()
	if err != nil {
		panic(err)
	}

}

//Murder murders the running process
func (p *Proc) Murder() {
	p.proc.Process.Kill()
	p.proc.Process.Wait()
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
