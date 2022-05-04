package command

import (
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/webhippie/cursecli/pkg/config"
	"github.com/webhippie/cursecli/pkg/forgesvc"
	"github.com/webhippie/cursecli/pkg/manifest"
)

// ManifestCmd provides the sub-command manifest.
func ManifestCmd(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:        "manifest",
		Usage:       "Manifest related commands",
		Flags:       ManifestFlags(cfg),
		Subcommands: ManifestCmds(cfg),
	}
}

// ManifestFlags defines the flags for the manifest sub-command.
func ManifestFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "manifest",
			Value:       "manifest.json",
			Usage:       "Path to manifest to parse",
			EnvVars:     []string{"CURSECLI_MANIFEST"},
			Destination: &cfg.General.Manifest,
		},
	}
}

// ManifestCmds defines the sub-commands for the manifest command.
func ManifestCmds(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		ManifestDownloadCmd(cfg),
	}
}

// ManifestDownloadCmd provides the sub-command manifest download.
func ManifestDownloadCmd(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:        "download",
		Usage:       "Download related commands",
		Flags:       ManifestDownloadFlags(cfg),
		Subcommands: ManifestDownloadCmds(cfg),
	}
}

// ManifestDownloadFlags defines the flags for the manifest download sub-command.
func ManifestDownloadFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{}
}

// ManifestDownloadCmds defines the sub-commands for the manifest download command.
func ManifestDownloadCmds(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		ManifestDownloadModsCmd(cfg),
	}
}

// ManifestDownloadModsCmd provides the sub-command manifest download mods.
func ManifestDownloadModsCmd(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "mods",
		Usage:  "Download mods defined within manifest",
		Flags:  ManifestDownloadModsFlags(cfg),
		Action: ManifestDownloadModsAction(cfg),
	}
}

// ManifestDownloadModsFlags defines the flags for the manifest download mods sub-command.
func ManifestDownloadModsFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Value:       "mods/",
			Usage:       "Path to download destination",
			EnvVars:     []string{"CURSECLI_PATH", "CURSECLI_DOWNLOAD_MODS_PATH"},
			Destination: &cfg.General.Path,
		},
	}
}

// ManifestDownloadModsAction implements the action for the manifest download mods command.
func ManifestDownloadModsAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		m, err1 := manifest.New(
			manifest.WithPath(c.String("manifest")),
		)

		if err1 != nil {
			log.Error().
				Err(err1).
				Msg("Failed to parse manifest")

			return err1
		}

		f, err2 := forgesvc.New(
			forgesvc.WithPath(c.String("path")),
			forgesvc.WithManifest(m),
		)

		if err2 != nil {
			log.Error().
				Err(err2).
				Msg("Failed to initalize client")

			return err2
		}

		if err := f.DownloadManifest(); err != nil {
			return err
		}

		return nil
	}
}
