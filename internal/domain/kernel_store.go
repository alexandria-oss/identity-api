package domain

import "github.com/spf13/viper"

type KernelStore struct {
	Service string
	Version string
	Config  struct {
		Cognito struct {
			UserPoolID        string
			UserPoolSecretKey string
		}
		Cache struct {
			Network  string
			Address  []string
			Username string
			Password string
			Database int
		}
	}
}

func init() {
	viper.SetDefault("alexandria.service", "identity")
	viper.SetDefault("alexandria.version", "1")
	viper.SetDefault("alexandria.aws.cognito.user-pool-id", "")
	viper.SetDefault("alexandria.aws.cognito.user-secret-key", "")
	viper.SetDefault("alexandria.persistence.cache.address", []string{"localhost:6379"})
	viper.SetDefault("alexandria.persistence.cache.network", "tcp")
	viper.SetDefault("alexandria.persistence.cache.username", "")
	viper.SetDefault("alexandria.persistence.cache.password", "")
	viper.SetDefault("alexandria.persistence.cache.database", 0)
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

	kernel.Service = viper.GetString("alexandria.service")
	kernel.Version = viper.GetString("alexandria.version")

	kernel.Config.Cognito.UserPoolID = viper.GetString("alexandria.aws.cognito.user-pool-id")
	kernel.Config.Cognito.UserPoolSecretKey = viper.GetString("alexandria.aws.cognito.user-secret-key")

	kernel.Config.Cache.Address = viper.GetStringSlice("alexandria.persistence.cache.address")
	kernel.Config.Cache.Network = viper.GetString("alexandria.persistence.cache.network")
	kernel.Config.Cache.Username = viper.GetString("alexandria.persistence.cache.username")
	kernel.Config.Cache.Password = viper.GetString("alexandria.persistence.cache.password")
	kernel.Config.Cache.Database = viper.GetInt("alexandria.persistence.cache.database")

	return kernel
}
