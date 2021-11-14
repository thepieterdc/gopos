package configuration

import "github.com/spf13/viper"

// Configuration wraps the configuration
type Configuration struct {
	config *viper.Viper
}
