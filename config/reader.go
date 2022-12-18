package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

// Read .
func Read(path string) (*Config, error) {
	var config Config
	var (
		dir      = filepath.Dir(path)
		filename = filepath.Base(path)
		ext      = filepath.Ext(path)
	)

	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext[1:])

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
