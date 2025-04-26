package hello

import (
	"github.com/aileron-gateway/aileron-gateway/apis/kernel"
	"github.com/aileron-gateway/aileron-gateway/core"
	"github.com/aileron-gateway/aileron-gateway/kernel/api"
	"github.com/aileron-gateway/aileron-gateway/kernel/log"
	utilhttp "github.com/aileron-gateway/aileron-gateway/util/http"
	v1 "github.com/aileron-gateway/example-extension/apis/sample/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	apiVersion = "ext/v1"
	kind       = "HelloHeaderMiddleware"
	Key        = apiVersion + "/" + kind
)

var (
	Resource api.Resource = &sampleAPI{&api.BaseResource{}}
)

type sampleAPI struct {
	*api.BaseResource
}

// Default must return default proto message instance.
// Metadata and Spec must not be nil.
func (*sampleAPI) Default() protoreflect.ProtoMessage {
	return &v1.HelloHeaderMiddleware{
		APIVersion: apiVersion,
		Kind:       kind,
		Metadata: &kernel.Metadata{
			Namespace: "default",
			Name:      "default",
		},
		Spec: &v1.HelloHeaderMiddlewareSpec{
			Value: "World !!",
		},
	}
}

// Create return a new instance of helloHeader middleware.
func (*sampleAPI) Create(a api.API[*api.Request, *api.Response], msg protoreflect.ProtoMessage) (any, error) {

	c := msg.(*v1.HelloHeaderMiddleware)

	// You can get an error handler for HTTP server.
	// Default error handler will be returned when c.Spec.ErrorHandler is nil.
	eh, err := utilhttp.ErrorHandler(a, c.Spec.ErrorHandler)
	if err != nil {
		return nil, core.ErrCoreGenCreateObject.WithStack(err, map[string]any{"kind": kind})
	}

	return &helloHeader{
		lg:    log.GlobalLogger(log.DefaultLoggerName),
		eh:    eh,
		value: c.Spec.Value,
	}, nil
}
