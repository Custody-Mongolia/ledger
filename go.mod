module github.com/numary/ledger

go 1.16

require (
	github.com/XSAM/otelsql v0.12.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/google/go-cmp v0.5.7
	github.com/huandu/go-sqlbuilder v1.13.0
	github.com/jackc/pgconn v1.10.1
	github.com/jackc/pgx/v4 v4.14.1
	github.com/mattn/go-sqlite3 v1.14.9
	github.com/numary/machine v1.1.0
	github.com/ory/dockertest/v3 v3.8.1
	github.com/pborman/uuid v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.28.0
	go.opentelemetry.io/otel v1.5.0
	go.opentelemetry.io/otel/exporters/jaeger v1.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v0.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.5.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.5.0
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/sdk v1.5.0
	go.opentelemetry.io/otel/sdk/metric v0.27.0
	go.opentelemetry.io/otel/trace v1.5.0
	go.uber.org/dig v1.14.1
	go.uber.org/fx v1.17.1
)

require (
	github.com/Shopify/sarama v1.32.0
	github.com/ThreeDotsLabs/watermill v1.1.1
	github.com/ThreeDotsLabs/watermill-http v1.1.4 // indirect
	github.com/ThreeDotsLabs/watermill-kafka/v2 v2.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-redis/redismock/v8 v8.0.6
	github.com/numary/go-libs v0.0.0-20220325080412-02604f34c81f
	github.com/xdg-go/scram v1.1.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
