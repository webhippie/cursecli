package command

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/cursecli/pkg/forge"
	"github.com/webhippie/cursecli/pkg/manifest"
)

var (
	manifestCmd = &cobra.Command{
		Use:   "manifest",
		Short: "Manifest related commands",
	}

	manifestDownloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Download related commands",
	}

	manifestDownloadModsCmd = &cobra.Command{
		Use:   "mods",
		Short: "Download mods defined within manifest",
		Run:   manifestDownloadModsAction,
	}

	defaultManifestPath = "manifest.json"
	defaultModsPath     = "mods/"
)

func init() {
	rootCmd.AddCommand(manifestCmd)
	manifestCmd.AddCommand(manifestDownloadCmd)
	manifestDownloadCmd.AddCommand(manifestDownloadModsCmd)

	manifestCmd.PersistentFlags().String("manifest", defaultManifestPath, "Path to manifest to parse")
	viper.SetDefault("manifest.path", defaultManifestPath)
	viper.BindPFlag("manifest.path", manifestCmd.PersistentFlags().Lookup("manifest"))

	manifestDownloadModsCmd.PersistentFlags().String("path", defaultModsPath, "Path to download destination")
	viper.SetDefault("mods.path", defaultModsPath)
	viper.BindPFlag("mods.path", manifestDownloadModsCmd.PersistentFlags().Lookup("path"))
}

func manifestDownloadModsAction(ccmd *cobra.Command, args []string) {
	m, err1 := manifest.New(
		manifest.WithPath(viper.GetString("manifest.path")),
	)

	if err1 != nil {
		log.Error().
			Err(err1).
			Msg("Failed to parse manifest")

		os.Exit(1)
	}

	f, err2 := forge.New(
		forge.WithPath(viper.GetString("mods.path")),
		forge.WithAPIKey(viper.GetString("api.key")),
		forge.WithManifest(m),
	)

	if err2 != nil {
		log.Error().
			Err(err2).
			Msg("Failed to initalize client")

		os.Exit(1)
	}

	if err := f.DownloadManifest(); err != nil {
		os.Exit(1)
	}
}
