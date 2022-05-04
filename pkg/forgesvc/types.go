package forgesvc

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
	ID     int    `json:"id"`
	Name   string `json:"fileName"`
	Size   int    `json:"fileLength"`
	URL    string `json:"downloadUrl"`
	Hashes []Hash `json:"hashes"`
}

// Hash represents a single hash from the forgesvc API.
type Hash struct {
	Algo  HashAlgo `json:"algorithm"`
	Value string   `json:"value"`
}
