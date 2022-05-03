package command

import (
	"github.com/urfave/cli/v2"
	"github.com/webhippie/cursecli/pkg/config"
)

// Download provides the sub-command to download modpacks.
func Download(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "download",
		Usage:  "Download modpack by manifest",
		Flags:  downloadFlags(cfg),
		Before: downloadBefore(cfg),
		Action: downloadAction(cfg),
	}
}

func downloadFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "manifest",
			Value:       "manifest.json",
			Usage:       "Path to manifest to parse",
			EnvVars:     []string{"CURSECLI_MANIFEST", "CURSECLI_DOWNLOAD_MANIFEST"},
			Destination: &cfg.General.Manifest,
		},
		&cli.StringFlag{
			Name:        "path",
			Value:       "",
			Usage:       "Path to download the mods into",
			EnvVars:     []string{"CURSECLI_PATH", "CURSECLI_DOWNLOAD_PATH"},
			Destination: &cfg.General.Path,
		},
	}
}

func downloadBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		return setupLogger(cfg)
	}
}

func downloadAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		return nil
	}
}
