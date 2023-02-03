package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

func loadFromYAML() ([]dataset, error) {
	// TODO fix get by args
	f, err := os.Open("tests/bqgen.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []dataset
	if err := yaml.NewDecoder(f).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// func loadBqDataFromConfig() ([]dataset, error) {
// 	con, err := loadConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var out []dataset
// 	if err := mapstructure.Decode(con, &out); err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func loadConfig() (map[string]any, error) {
// 	// configHome := os.Getenv("XDG_CONFIG_HOME")
// 	// homePath := os.Getenv("HOME")
// 	// wd, err := os.Getwd()
// 	// if err != nil {
// 	// 	wd = "."
// 	// }

// 	// configPaths := []string{wd}
// 	// if len(configHome) > 0 {
// 	// 	configPaths = append(configPaths, filepath.Join(configHome, "scheman"))
// 	// } else {
// 	// 	configPaths = append(configPaths, filepath.Join(homePath, ".config/scheman"))
// 	// }

// 	// for _, p := range configPaths {
// 	// 	viper.AddConfigPath(p)
// 	// }

// 	viper.AddConfigPath(".")
// 	viper.SetConfigFile("bqgen")
// 	if err := viper.ReadInConfig(); err != nil {
// 		// fmt.Println(err)
// 	}

// 	return viper.AllSettings(), nil
// }
