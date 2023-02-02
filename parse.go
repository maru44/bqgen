package main

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func loadBqDataFromConfig() ([]dataset, error) {
	con, err := loadConfig()
	if err != nil {
		return nil, err
	}
	var out []dataset
	if err := mapstructure.Decode(con, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func loadConfig() (map[string]any, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	homePath := os.Getenv("HOME")
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	configPaths := []string{wd}
	if len(configHome) > 0 {
		configPaths = append(configPaths, filepath.Join(configHome, "scheman"))
	} else {
		configPaths = append(configPaths, filepath.Join(homePath, ".config/scheman"))
	}

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	// Ignore errors here, fallback to other validation methods.
	// Users can use environment variables if a config is not found.
	_ = viper.ReadInConfig()

	viper.AddConfigPath(".")
	viper.SetConfigFile("bqgen")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.AllSettings(), nil
}
