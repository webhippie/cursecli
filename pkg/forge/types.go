package forge

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// HashAlgo defines the available hash algorithms.
type HashAlgo int

const (
	// HashAlgoSha1 defines the SHA1 hash algorithm.
	HashAlgoSha1 HashAlgo = 1

	// HashAlgoMd5 defines the MD5 hash algorithm.
	HashAlgoMd5 HashAlgo = 2
)

func (h HashAlgo) String() string {
	return HashAlgoToString[h]
}

// HashAlgoToString converts the hash algorithm to string.
var HashAlgoToString = map[HashAlgo]string{
	HashAlgoSha1: "SHA1",
	HashAlgoMd5:  "MD5",
}

// HashAlgoToID converts the hash algorithm to integer.
var HashAlgoToID = map[string]HashAlgo{
	"SHA1": HashAlgoSha1,
	"MD5":  HashAlgoMd5,
}

// File represents a single file from the forgesvc API.
type File struct {
	ID       int    `json:"id"`
	Name     string `json:"fileName"`
	Size     int    `json:"fileLength"`
	URL      string `json:"downloadUrl"`
	Fallback bool   `json:"-"`
	Hashes   []Hash `json:"hashes"`
}

// UnmarshalJSON implements the JSON unmarshaling.
func (f *File) UnmarshalJSON(b []byte) error {
	type F File
	var v F

	if err := json.Unmarshal(b, &struct{ Data *F }{&v}); err != nil {
		return err
	}

	*f = File(v)

	if f.URL == "" {
		identifier := strconv.Itoa(f.ID)

		f.URL = fmt.Sprintf(
			"https://edge.forgecdn.net/files/%s/4%s/%s",
			identifier[0:4],
			identifier[5:len(identifier)-1],
			f.Name,
		)

		f.Fallback = true
	}

	return nil
}

// Hash represents a single hash from the forgesvc API.
type Hash struct {
	Algo  HashAlgo `json:"algo"`
	Value string   `json:"value"`
}
