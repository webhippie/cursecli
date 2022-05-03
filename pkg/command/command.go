package command

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/webhippie/cursecli/pkg/config"
	"github.com/webhippie/cursecli/pkg/version"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.App{
		Name:     "cursecli",
		Version:  version.String,
		Usage:    "Commandline client for Curseforge",
		Authors:  authorList(),
		Flags:    globalFlags(cfg),
		Commands: globalCommands(cfg),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	return app.Run(os.Args)
}

func authorList() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Thomas Boerger",
			Email: "thomas@webhippie.de",
		},
	}
}

func globalFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVars:     []string{"CURSECLI_LOG_LEVEL"},
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log-pretty",
			Value:       true,
			Usage:       "Enable pretty logging",
			EnvVars:     []string{"CURSECLI_LOG_PRETTY"},
			Destination: &cfg.Logs.Pretty,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Value:       true,
			Usage:       "Enable colored logging",
			EnvVars:     []string{"CURSECLI_LOG_COLOR"},
			Destination: &cfg.Logs.Color,
		},
	}
}

func globalCommands(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		Download(cfg),
	}
}
