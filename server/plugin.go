package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ericjaystevens/slashparse"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

const configPath = "/home/ec2-user/code/slashparse_sample1/server/simple.yaml"
// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
	slashCommand slashparse.SlashCommand
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) OnActivate() error {
	if err := p.API.RegisterCommand(getCommand()); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to register command: %v", getCommand()))
	}
	slashDef, _ := ioutil.ReadFile(configPath)

	p.slashCommand, _ = slashparse.NewSlashCommand(slashDef)
	
	return nil
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
