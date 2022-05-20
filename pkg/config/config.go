package config

// Manifest defines the manifest stuff like path.
type Manifest struct {
	Path string
}

// Mods defines the mods stuff like path.
type Mods struct {
	Path string
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
	Color  bool
}

// Config is a combination of all available configurations.
type Config struct {
	Manifest Manifest
	Mods     Mods
	Logs     Logs
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
