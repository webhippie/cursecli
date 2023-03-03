package manifest

import (
	"encoding/json"
	"fmt"
	"os"
)

// Manifest defines the manifest itself.
type Manifest struct {
	Files []File `json:"files"`
}

// File defines a record within manifest.
type File struct {
	ProjectID int  `json:"projectID"`
	FileID    int  `json:"fileID"`
	Required  bool `json:"required"`
}

// New parses and prepares a manifest definition.
func New(opts ...Option) (Manifest, error) {
	sopts := newOptions(opts...)
	manifest := Manifest{}

	if _, err := os.Stat(sopts.Path); err != nil && os.IsNotExist(err) {
		return manifest, fmt.Errorf("manifest does not exist: %w", err)
	} else if err != nil {
		return manifest, fmt.Errorf("failed to check manifest: %w", err)
	}

	rawdata, err := os.ReadFile(sopts.Path)

	if err != nil {
		return manifest, fmt.Errorf("failed to read manifest: %w", err)
	}

	if err := json.Unmarshal(rawdata, &manifest); err != nil {
		return manifest, fmt.Errorf("failed to parse manifest: %w", err)
	}

	return manifest, nil
}
