package log

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/patex-ecosystem/patex-network/pt-service/httputil"
)

func NewLoggingMiddleware(lgr log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := httputil.NewWrappedResponseWriter(w)
		start := time.Now()
		next.ServeHTTP(ww, r)
		lgr.Debug(
			"served HTTP request",
			"status", ww.StatusCode,
			"response_len", ww.ResponseLen,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
			"remote_addr", r.RemoteAddr,
		)
	})
}
