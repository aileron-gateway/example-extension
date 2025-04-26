package hello

import (
	"errors"
	"net/http"

	"github.com/aileron-gateway/aileron-gateway/core"
	"github.com/aileron-gateway/aileron-gateway/kernel/log"
	utilhttp "github.com/aileron-gateway/aileron-gateway/util/http"
)

// Check that the [helloHeader] implements
// [core.Middleware] interface.
var _ core.Middleware = &helloHeader{}

type helloHeader struct {
	lg    log.Logger
	eh    core.ErrorHandler
	value string
}

func (m *helloHeader) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Example of log output.
		if m.lg.Enabled(log.LvDebug) {
			m.lg.Debug(r.Context(), r.Method+r.URL.Path)
		}

		// Example of error handling.
		if r.URL.Path == "/error" {
			err := errors.New("hello header middleware error")
			m.eh.ServeHTTPError(w, r, utilhttp.NewHTTPError(err, http.StatusInternalServerError))
			return
		}

		// Add "Hello" header to the response.
		w.Header().Set("Hello", m.value)

		// Call next middleware or handler.
		next.ServeHTTP(w, r)

	})
}
