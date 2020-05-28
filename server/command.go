package main

import (
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

	split := strings.Fields(args.Command)
	newSlash, err := slashparse.NewSlashCommand(split, commandDefPath)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	command, err := newSlash.GetCommandString(split)
	if err != nil {
		return &model.CommandResponse{Text: err.Error()}, nil
	}

	switch command {
	case "print help":
		text := newSlash.GetSlashHelp()
		return &model.CommandResponse{Text: text}, nil
	case "print":
		text := executePrint()
		return &model.CommandResponse{Text: text}, nil
	default:
		text := "Unknown unknown"
		return &model.CommandResponse{Text: text}, nil
	}
}

func executePrint() (msg string) {
	msg = "you want my to say what?"
	return
}
