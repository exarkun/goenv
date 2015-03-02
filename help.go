package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/template"
)

var helpCommand = Command{
	Name:    "help",
	Short:   "get help for a command",
	Usage:   "help [command]",
	Long:    "TODO",
	GetTask: NewHelpTask,
}

var helpTemplate = `
Usage: goenv {{.Usage}}

{{.Long}}
`

// HelpTask initializes a goenv.
type HelpTask struct {
	CommandName string              // the name of the command.
	Commands    map[string]*Command // the map of commands.
}

// NewHelpTask returns a new HelpTask created with the specified command-line args.
func NewHelpTask(args []string) (Task, error) {

	flags := flag.NewFlagSet("help", flag.ExitOnError)
	flags.Parse(args)
	args = flags.Args()

	if len(args) < 1 {
		return nil, errors.New("no command specified")
	}

	return &HelpTask{
		CommandName: args[0],
		Commands:    commands,
	}, nil
}

// Run exeuctes the HelpTask
func (task *HelpTask) Run() error {

	cmd, found := task.Commands[task.CommandName]

	if !found {
		return fmt.Errorf("no such command \"%s\"", cmd)
	}

	tmpl := template.New("help")
	tmpl, err := tmpl.Parse(helpTemplate)

	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, cmd)

	return err
}
