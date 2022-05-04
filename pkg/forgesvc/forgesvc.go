package forgesvc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/webhippie/cursecli/pkg/manifest"
)

const (
	downloadURL = "https://addons-ecs.forgesvc.net/api/v2/addon/%d/file/%d"
)

// Forgesvc defines the forgesvc itself.
type Forgesvc struct {
	HTTPClient *http.Client
	Path       string
	Manifest   manifest.Manifest
}

// New parses and prepares a manifest definition.
func New(opts ...Option) (*Forgesvc, error) {
	sopts := newOptions(opts...)

	return &Forgesvc{
		HTTPClient: sopts.HTTPClient,
		Path:       sopts.Path,
		Manifest:   sopts.Manifest,
	}, nil
}

// DownloadManifest downloads all mods defined within a manifest.
func (f *Forgesvc) DownloadManifest() error {
	if err := os.MkdirAll(f.Path, os.ModePerm); err != nil {
		log.Error().
			Err(err).
			Str("path", f.Path).
			Msg("Failed to create mod directory")

		return err
	}

	for _, file := range f.Manifest.Files {
		resp, err := f.HTTPClient.Get(
			fmt.Sprintf(
				downloadURL,
				file.ProjectID,
				file.FileID,
			),
		)

		if err != nil {
			log.Error().
				Err(err).
				Int("project", file.ProjectID).
				Int("file", file.FileID).
				Msg("Failed to fetch mod details")

			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Error().
				Err(err).
				Int("project", file.ProjectID).
				Int("file", file.FileID).
				Msg("Failed to read mod details")

			return err
		}

		download := File{}

		if err := json.Unmarshal(body, &download); err != nil {
			log.Error().
				Err(err).
				Int("project", file.ProjectID).
				Int("file", file.FileID).
				Msg("Failed to parse mod details")

			return err
		}

		if err := f.downloadFile(
			download.Name,
			download.URL,
		); err != nil {
			return err
		}

		log.Info().
			Int("project", file.ProjectID).
			Int("file", file.FileID).
			Str("name", download.Name).
			Msg("Successfully downloaded mod")
	}

	return nil
}

func (f *Forgesvc) downloadFile(name, url string) error {
	resp, err := f.HTTPClient.Get(
		url,
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	out, err := os.Create(filepath.Join(f.Path, name))

	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	return err
}
