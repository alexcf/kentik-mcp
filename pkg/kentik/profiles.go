package kentik

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ProfileConfig is the format of ~/.kentik-mcp.json.
type ProfileConfig struct {
	Profiles []ProfileEntry `json:"profiles"`
}

// ProfileEntry is one entry in the profiles config file.
type ProfileEntry struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
	Region   string `json:"region"`
}

// DefaultConfigPath returns the default path for the profiles config file.
func DefaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".kentik-mcp.json")
}

// LoadProfileConfig reads the profiles config file. Returns an empty config
// (not an error) if the file does not exist.
func LoadProfileConfig(path string) (*ProfileConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &ProfileConfig{}, nil
		}
		return nil, fmt.Errorf("read profile config %s: %w", path, err)
	}
	var cfg ProfileConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse profile config %s: %w", path, err)
	}
	return &cfg, nil
}
