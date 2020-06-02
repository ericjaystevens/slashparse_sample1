package main

import (
	"io/ioutil"
	"strings"

	"github.com/ericjaystevens/slashparse"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:     "print",
		DisplayName: "Print",
	}
}

const commandDefPath = "/home/ec2-user/code/slashparse_sample1/server/simple.yaml"

// ExecuteCommand is entrypoint for print command
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	commandDef, err := ioutil.ReadFile(commandDefPath)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	newSlash, err := slashparse.NewSlashCommand(args.Command, commandDef)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	command, err := newSlash.GetCommandString(args.Command)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	arguments, err := newSlash.GetValues(args.Command)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	switch command {
	case "Print":
		text := executePrint(arguments, &newSlash) //this newSlash will go away after subcommand allow help call in case statement
		return &model.CommandResponse{Text: text}, nil
	default:
		text := "Unknown unknown"
		return &model.CommandResponse{Text: text}, nil
	}
}

func executePrint(values map[string]string, newSlash *slashparse.SlashCommand) (msg string) {

	if strings.EqualFold(values["text"], "help") {
		msg = newSlash.GetSlashHelp()
	} else {
		msg = "you want my to say what? ...  " + values["text"]
	}
	return
}
