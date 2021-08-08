package configuration

import (
	"fmt"

	"github.com/GolangTechTask/pkg/constant"
	"github.com/spf13/viper"
)

const (
	envPrefix = "VT"
)

//VotingConfiguration configuration map with default value
var VotingConfiguration = map[string]interface{}{
	constant.LocalGrpcPort: 9090,
	constant.DbEndPoint:    "http://localhost:8000",
	constant.AWSRegion:     "us-west-2",
	//AWSKey ... set AWSKey in ENV variable
	constant.AWSKey:        "",
	constant.AWSSecret:     "",
	constant.LogTimeFormat: "2006-01-02T15:04:05.99",
	constant.LogLevel:      -1,
}

// LoadDefaults loads default values from vanguardConfiguration in the configuration manager
func LoadDefaults() {
	LoadConfigDefaults(VotingConfiguration, envPrefix)
}

//LoadConfigDefaults  ...
func LoadConfigDefaults(defaults map[string]interface{}, envPrefix string) {
	for key, value := range defaults {
		if value != nil {
			viper.SetDefault(key, value)
		}
	}

	if envPrefix != "" {
		viper.SetEnvPrefix(envPrefix)
		viper.AutomaticEnv()
	}
}

//RequireString ...
func RequireString(key string) string {
	assertSet(key)
	return viper.GetString(key)
}

//RequireInt ...
func RequireInt(key string) int {
	assertSet(key)
	return viper.GetInt(key)
}

func assertSet(key string) {
	if !viper.IsSet(key) {
		handleNotSetError(key)
	}
}

func handleNotSetError(key string) {
	panic(fmtNotSetError(key).Error())
}

func fmtNotSetError(key string) error {
	return fmt.Errorf("required configuration parameter '%s' not set", key)
}
