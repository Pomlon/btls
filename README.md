# btls
Simple build tools, can watch directories for changes, ignore certain files/folders and restart processes. Configurable by json

btls can be used for any programming language, since the commands from config are executed via exec.Command()

## Why was this made? Why use this?
I wrote this because I needed a watch+build tool for a project of mine. What irritated me with other implementations was that
they tended to sometimes not register code changes, fail restarts and loose output from the watched executable.

This util so far resolved all of my problems, but be wary! It outright KILLS the watched process which results in a faster restart,
but may reult in corrupted data here and there.

Another strength of this tool is that it isn't only a watcher, it can also execute any command you input into the config, feel
free to experiment.

## Installation

    go get -u github.com/Pomlon/btls
    go install github.com/Pomlon/btls
  
# Usage
buildConfTemplate.json includes a full bconf template.

TaskName : name your task

TaskDescr : Describe it as you wish, use \n for multiline descriptions

Build : build command(s)

AfterBuild : command(s) executed after build, this is mostly cosmetic

WatchPath : path to be watched for changes (recursive)

IgnorePaths : list of things to ignore while watching the path.

RunCommand : If this is present it will be ran with stdio and stderr redirected to console and restarted when a change is detected

## Example config

      [
        {
          "TaskName" : "watch",
          "TaskDescr" : "Watch project for changes, build and run server",
          "Build" : ["go build"],
          "AfterBuild" : [],
          "WatchPath" : ".",
          "IgnorePaths" : [".gitignore", "./static", "./templates", "./release", "./App.exe", "./test.db", "./.git"],
          "RunCommand" : "App.exe"
        }
      ]
      
This config builds a go app I'm working on and restarts it's server. Ignores stuff like git, templates, static and an SQLite test db.

# Available commands

    btls list - list defined tasks and descriptions, also no-param run will do the same

    btls <taskName> - execute given task
  
    btls <taskName> help - displays task name + extended description

## Notes

* Util panics if no buildConfig is present.
* Undefined behavior on invalid configs, most likely non-malicious though
