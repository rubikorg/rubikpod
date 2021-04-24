package app

// PodServerConfig defines the release config properties
type PodServerConfig struct {
	Path string `toml:"path"`
}

type AuthConfig struct {
	Access   []string `yaml:"access"`
	Password string   `yaml:"password"`
}

type ProjectDetails struct {
	Path    string
	Cmd     string
	GitBase string `yaml:"git_base"`
	Branch  string `yaml:"branch"`
}

type GitConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type PodDetails struct {
	Git      map[string]GitConfig      `yaml:"git"`
	Users    map[string]AuthConfig     `yaml:"users"`
	Projects map[string]ProjectDetails `yaml:"projects"`
}

// ProjectConfig defines your server related configuration
type ProjectConfig struct {
	Port            int             `toml:"port"`
	PodServerConfig PodServerConfig `toml:"pod_server"`
	PodDetails      PodDetails
}
