package common

import "github.com/spf13/viper"

type KernelStore struct {
	Service string
	Version string
	Config  configStore
}

type configStore struct {
	Cognito struct {
		UserPoolID        string
		UserPoolSecretKey string
	}
}

func init() {
	viper.SetDefault("alexandria.service", "identity")
	viper.SetDefault("alexandria.version", "1")
	viper.SetDefault("alexandria.aws.cognito.user-pool-id", "")
	viper.SetDefault("alexandria.aws.cognito.user-secret-key", "")
}

func NewKernelStore() KernelStore {
	kernel := KernelStore{}

	viper.SetConfigName("alexandria-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	// Open config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			_ = viper.SafeWriteConfig()
		}

		// Config file was found but another error was produced, use default values
	}

	kernel.Config.Cognito.UserPoolID = viper.GetString("alexandria.aws.cognito.user-pool-id")
	kernel.Config.Cognito.UserPoolSecretKey = viper.GetString("alexandria.aws.cognito.user-secret-key")

	return kernel
}
