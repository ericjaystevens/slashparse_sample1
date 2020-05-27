package main

import (
	//"strings"

	//"github.com/ericjaystevens/slashparse"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:     "print",
		DisplayName: "Print",
	}
}

const commandDefPath = "./simple.yaml"

// ExecuteCommand is entrypoint for print command
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	//split := strings.Fields(args.Command)
	//newSlash, err := slashparse.NewSlashCommand(split, commandDefPath)
	//if err != nil {
		//return &model.CommandResponse{Text: "here"}, &model.AppError{Message: "unable to process yaml"}
//	}

	//_, err = newSlash.GetCommandString(split)
	//if err != nil {
		//return &model.CommandResponse{Text: "actually here too"}, &model.AppError{Message: "unable to determine command"}
	//}

	msg := "I must print things now, when you enter:?"
	response := &model.CommandResponse{
		Text: msg,
	}

	return response, nil
}
