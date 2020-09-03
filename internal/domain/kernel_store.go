// Copyright 2020 The Alexandria Foundation
//
// Licensed under the GNU Affero General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"github.com/spf13/viper"
	"strings"
)

type KernelStore struct {
	Service     string
	Version     string
	Environment string
	APIVersion  string
	Config      struct {
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
		Transport struct {
			HTTP struct {
				Address string
			}
			GRPC struct {
				Address string
			}
		}
		Tracing struct {
			Collector string
			Agent     string
			Username  string
			Password  string
		}
	}
}

func init() {
	viper.SetDefault("alexandria.service", "identity")
	viper.SetDefault("alexandria.version", "1")
	viper.SetDefault("alexandria.environment", "dev")
	viper.SetDefault("alexandria.api-version", "v1")
	viper.SetDefault("alexandria.aws.cognito.user-pool-id", "")
	viper.SetDefault("alexandria.aws.cognito.user-secret-key", "")
	viper.SetDefault("alexandria.persistence.cache.address", []string{"localhost:6379"})
	viper.SetDefault("alexandria.persistence.cache.network", "tcp")
	viper.SetDefault("alexandria.persistence.cache.username", "")
	viper.SetDefault("alexandria.persistence.cache.password", "")
	viper.SetDefault("alexandria.persistence.cache.database", 0)
	viper.SetDefault("alexandria.tracing.jaeger.collector", "http://localhost:14268/api/traces")
	viper.SetDefault("alexandria.tracing.jaeger.agent", "localhost:6831")
	viper.SetDefault("alexandria.tracing.jaeger.username", "")
	viper.SetDefault("alexandria.tracing.jaeger.password", "")
	viper.SetDefault("alexandria.transport.http.address", "localhost:8080")
	viper.SetDefault("alexandria.transport.grpc.address", "localhost:9090")
}

func NewKernelStore() KernelStore {
	kernel := KernelStore{}

	viper.SetConfigName("alexandria-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")

	// Open config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			_ = viper.SafeWriteConfig()
		}

		// Config file was found but another error was produced, use default values
	}

	kernel.Service = strings.ToLower(viper.GetString("alexandria.service"))
	kernel.Version = strings.ToLower(viper.GetString("alexandria.version"))
	kernel.Environment = SanitizeEnvironment(strings.ToLower(viper.GetString("alexandria.environment")))
	kernel.APIVersion = strings.ToLower(viper.GetString("alexandria.api-version"))

	kernel.Config.Cognito.UserPoolID = viper.GetString("alexandria.aws.cognito.user-pool-id")
	kernel.Config.Cognito.UserPoolSecretKey = viper.GetString("alexandria.aws.cognito.user-secret-key")

	kernel.Config.Cache.Address = viper.GetStringSlice("alexandria.persistence.cache.address")
	kernel.Config.Cache.Network = viper.GetString("alexandria.persistence.cache.network")
	kernel.Config.Cache.Username = viper.GetString("alexandria.persistence.cache.username")
	kernel.Config.Cache.Password = viper.GetString("alexandria.persistence.cache.password")
	kernel.Config.Cache.Database = viper.GetInt("alexandria.persistence.cache.database")
	kernel.Config.Tracing.Collector = viper.GetString("alexandria.tracing.jaeger.collector")
	kernel.Config.Tracing.Agent = viper.GetString("alexandria.tracing.jaeger.agent")
	kernel.Config.Tracing.Username = viper.GetString("alexandria.tracing.jaeger.username")
	kernel.Config.Tracing.Password = viper.GetString("alexandria.tracing.jaeger.password")
	kernel.Config.Transport.HTTP.Address = viper.GetString("alexandria.transport.http.address")
	kernel.Config.Transport.GRPC.Address = viper.GetString("alexandria.transport.grpc.address")

	return kernel
}
