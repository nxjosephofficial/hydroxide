package config

import (
	"os"
	"path/filepath"
)

func Path(filename string) (string, error) {
	configHome := "/var/db/hydroxide"

	p := filepath.Join(configHome, filename)

	dirname, _ := filepath.Split(p)
	if err := os.MkdirAll(dirname, 0700); err != nil {
		return "", err
	}

	return p, nil
}
