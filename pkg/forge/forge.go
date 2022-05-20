package forge

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
	downloadURL = "https://api.curseforge.com/v1/mods/%d/files/%d"
)

// Forge defines the forgesvc itself.
type Forge struct {
	HTTPClient *http.Client
	Path       string
	APIKey     string
	Manifest   manifest.Manifest
}

// New parses and prepares a manifest definition.
func New(opts ...Option) (*Forge, error) {
	sopts := newOptions(opts...)

	return &Forge{
		HTTPClient: sopts.HTTPClient,
		Path:       sopts.Path,
		APIKey:     sopts.APIKey,
		Manifest:   sopts.Manifest,
	}, nil
}

// DownloadManifest downloads all mods defined within a manifest.
func (f *Forge) DownloadManifest() error {
	if err := os.MkdirAll(f.Path, os.ModePerm); err != nil {
		log.Error().
			Err(err).
			Str("path", f.Path).
			Msg("Failed to create mod directory")

		return err
	}

	for _, file := range f.Manifest.Files {
		resp, err := f.get(
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
			fmt.Printf(
				downloadURL,
				file.ProjectID,
				file.FileID,
			)

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
			log.Error().
				Err(err).
				Int("project", file.ProjectID).
				Int("file", file.FileID).
				Str("name", download.Name).
				Msg("Failed to download mod")

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

func (f *Forge) downloadFile(name, url string) error {
	resp, err := f.get(
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

func (f *Forge) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return f.do(req)
}

func (f *Forge) do(req *http.Request) (*http.Response, error) {
	req.Header.Set(
		"x-api-key",
		f.APIKey,
	)

	return f.HTTPClient.Do(req)
}
