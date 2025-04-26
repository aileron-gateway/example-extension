package main

import (
	"fmt"
	"os"

	"github.com/aileron-gateway/aileron-gateway/app/handler/echo"
	"github.com/aileron-gateway/aileron-gateway/app/middleware/compression"
	"github.com/aileron-gateway/aileron-gateway/app/middleware/header"
	"github.com/aileron-gateway/aileron-gateway/app/middleware/throttle"
	"github.com/aileron-gateway/aileron-gateway/app/middleware/timeout"
	"github.com/aileron-gateway/aileron-gateway/app/middleware/tracking"
	"github.com/aileron-gateway/aileron-gateway/cmd/aileron/app"
	"github.com/aileron-gateway/aileron-gateway/core/entrypoint"
	"github.com/aileron-gateway/aileron-gateway/core/errhandler"
	"github.com/aileron-gateway/aileron-gateway/core/httpclient"
	"github.com/aileron-gateway/aileron-gateway/core/httphandler"
	"github.com/aileron-gateway/aileron-gateway/core/httplogger"
	"github.com/aileron-gateway/aileron-gateway/core/httpproxy"
	"github.com/aileron-gateway/aileron-gateway/core/httpserver"
	"github.com/aileron-gateway/aileron-gateway/core/slogger"
	"github.com/aileron-gateway/aileron-gateway/core/static"
	"github.com/aileron-gateway/aileron-gateway/core/template"
	"github.com/aileron-gateway/aileron-gateway/kernel/api"
	"github.com/aileron-gateway/example-extension/feature/hello"
)

func main() {
	a := app.New() // Create new app.
	a.ParseArgs(os.Args[1:])

	// Register features we want to use.
	f := api.NewFactoryAPI()
	f.Register(entrypoint.Key, entrypoint.Resource)
	f.Register(errhandler.Key, errhandler.Resource)
	f.Register(httpclient.Key, httpclient.Resource)
	f.Register(httphandler.Key, httphandler.Resource)
	f.Register(httplogger.Key, httplogger.Resource)
	f.Register(httpproxy.Key, httpproxy.Resource)
	f.Register(httpserver.Key, httpserver.Resource)
	f.Register(slogger.Key, slogger.Resource)
	f.Register(static.Key, static.Resource)
	f.Register(template.Key, template.Resource)
	f.Register(compression.Key, compression.Resource)
	f.Register(echo.Key, echo.Resource)
	f.Register(header.Key, header.Resource)
	f.Register(throttle.Key, throttle.Resource)
	f.Register(timeout.Key, timeout.Resource)
	f.Register(tracking.Key, tracking.Resource)
	// Register extensional features.
	f.Register(hello.Key, hello.Resource)

	server := api.NewDefaultServeMux()
	server.Handle("core/", f) // Handle core/v1 features.
	server.Handle("app/", f)  // Handle app/v1 features.
	server.Handle("ext/", f)  // Handle ext/v1 features.

	if err := a.Run(server); err != nil {
		e := app.ErrAppMain.WithStack(err, nil)
		fmt.Println(e.Error())
		fmt.Println(e.StackTrace())
		os.Exit(1)
	}
}
