package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:     "print",
		DisplayName: "Print",
	}
}

// ExecuteCommand is entrypoint for print command
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	command, values, err := p.slashCommand.Parse(args.Command)
	if err != nil {
		text := "bad command deffinition"
		return &model.CommandResponse{Text: text}, nil
	}


	switch command {
	case "Print":
		text := executePrint(values) //this newSlash will go away after subcommand allow help call in case statement
		return &model.CommandResponse{Text: text}, nil
	case "help":
		text := p.slashCommand.GetSlashHelp()
		return &model.CommandResponse{Text: text}, nil
	default:
		text := "Unknown unknown"
		return &model.CommandResponse{Text: text}, nil
	}
}

func executePrint(values map[string]string ) (msg string) {

	msg = "you want my to say what? ...  " + values["text"]
	return
}
