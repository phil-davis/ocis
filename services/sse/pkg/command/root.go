package command

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/owncloud/ocis/v2/ocis-pkg/clihelper"
	"github.com/owncloud/ocis/v2/services/sse/pkg/config"
)

// GetCommands provides all commands for this service
func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		Server(cfg),
		Health(cfg),
		Version(cfg),
	}
}

// Execute is the entry point for the sse command.
func Execute(cfg *config.Config) error {
	app := clihelper.DefaultApp(&cli.App{
		Name:     "sse",
		Usage:    "Serve ownCloud sse for oCIS",
		Commands: GetCommands(cfg),
	})

	return app.Run(os.Args)
}
