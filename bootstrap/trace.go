package bootstrap

import (
	"fmt"
	"hzminioupload/biz/pkg/global"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", global.S_CONFIG.GetString("trace.jaeger.agent..host"), global.S_CONFIG.GetInt("jaeger.agent.port")),
			CollectorEndpoint:  global.S_CONFIG.GetString("trace.jaeger.endpoint"),
		},
		ServiceName: serviceName,
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	// opentracing.InitGlobalTracer(tracer)
	return tracer, closer
}
