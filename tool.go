package main

import (
	"fmt"
	"os"
)

func main() {
	/*proc := Proc{}
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
	}*/

	tasks := ParseConfig()
	var runningTaskIndex int

	if len(os.Args) == 1 || os.Args[1] == "list" {
		fmt.Println()
		for _, task := range tasks {
			fmt.Println(task.TaskName)
			fmt.Println()
			fmt.Println(task.TaskDescr)
			fmt.Println("----------------------------------------------")
		}
		os.Exit(0)
	}

	for index, task := range tasks {
		if task.TaskName == os.Args[1] {
			runningTaskIndex = index
		}

	}

	runningTask := tasks[runningTaskIndex]

	if len(os.Args) > 2 && os.Args[2] == "help" {
		fmt.Println()
		fmt.Println("TASK: " + runningTask.TaskName)
		fmt.Println()
		fmt.Println(runningTask.TaskDescr)
		fmt.Println()
		fmt.Print("Build: ")
		fmt.Println(runningTask.Build)
		fmt.Println()
		fmt.Print("After build: ")
		fmt.Println(runningTask.AfterBuild)
		fmt.Println()
		fmt.Print("Path watch: ")
		fmt.Println(runningTask.WatchPath)
		fmt.Println()
		fmt.Print("Run command: ")
		fmt.Println(runningTask.RunCommand)
		os.Exit(0)
	}

	for {
		fmt.Println("Running build commands...")
		OneRunMany(runningTask.Build...)
		fmt.Println("Done!")

		fmt.Println("Running AfterBuild commands...")
		OneRunMany(runningTask.AfterBuild...)
		fmt.Println("Done!")

		if runningTask.WatchPath != "" {
			fmt.Println("Watching path...")
		} else {
			fmt.Println("Nothing to watch! Exiting...")
			os.Exit(0)
		}

	}
}
