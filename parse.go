package main

import (
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
	viper.AddConfigPath(".")
	viper.SetConfigFile("bqgen")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.AllSettings(), nil
}
