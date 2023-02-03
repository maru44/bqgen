package input

import (
	"context"

	"github.com/jellydator/validation"
	"github.com/spf13/viper"
)

type Config struct {
	// Dir is the directory possesses type definition files.
	// If `Files` length is not 0, `Files` is used for load definition.
	Dir string `json:"dir" yaml:"dir" toml:"dir"`
	// Files is the files written definition.
	Files []string `json:"files" yaml:"files" toml:"files"`
	// OutputPkg is where files are generated.
	Output string `json:"output" yaml:"output" toml:"output"`
	// Pkg is the package name for generated files.
	Pkg string `json:"pkg" yaml:"pkg" toml:"pkg"`
	// TfDir is the where terraform files are generated.
	TfDir string `json:"tfdir" yaml:"tfdir" toml:"tfdir"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("bqgen")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, c,
		validation.Field(&c.Output, validation.Required),
		validation.Field(&c.Pkg, validation.Required),
		validation.Field(&c.Dir, validation.When(len(c.Files) == 0, validation.Required)),
	)
}
