package config

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// Viper will check for an environment variable any time a viper.Get request is made.
	viper.AutomaticEnv()

	// get current file path
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	// get Test variable
	test := viper.GetString("TEST")
	var configPath string

	configPath = dir + "/../../../config/config.yml"
	if test == "True" {
		configPath = dir + "/../../../config/config.test.yml"
	}

	err := Parse(configPath)
	if err != nil {
		panic(err)
	}
}

// Parse parse config file
func Parse(path string) error {
	configType := strings.Replace(filepath.Ext(path), ".", "", 1)
	viper.SetConfigType(strings.ToLower(configType))
	viper.SetConfigFile(path)
	viper.WatchConfig()

	err := viper.ReadInConfig()
	return err
}

// Get get config value by key
func Get(key string) interface{} {
	return viper.Get(key)
}
