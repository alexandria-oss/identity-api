module github.com/alexandria-oss/identity-api

go 1.15

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.2.0
	github.com/alexandria-oss/common-go v0.1.0-alpha
	github.com/aws/aws-sdk-go v1.34.2
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.4.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	go.opencensus.io v0.22.4
	go.uber.org/ratelimit v0.1.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
