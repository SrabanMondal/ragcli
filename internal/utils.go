package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ashank007/docai/reader"
	"gopkg.in/yaml.v3"
)

func getReaderByExtension(path string) (reader.Reader, error) {
	switch filepath.Ext(path) {
	case ".pdf":
		return reader.NewPDFReader(), nil
	case ".txt":
		return reader.NewTextReader(), nil
	case ".docx":
		return reader.NewDocxReader(), nil
	default:
		return nil, fmt.Errorf("unsupported extension: %s", path)
	}
}

func SaveConfig(cfg *Config) error {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".ragcli")
	os.MkdirAll(path, 0755)

	filePath := filepath.Join(path, "config.yaml")
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func LoadConfig() *Config {
	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, ".ragcli", "config.yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
		// fallback to default
		return DefaultConfig()
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return DefaultConfig()
	}
	return &cfg
}

