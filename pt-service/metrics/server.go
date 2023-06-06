package metrics

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/patex-ecosystem/patex-network/pt-service/httputil"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ListenAndServe(ctx context.Context, r *prometheus.Registry, hostname string, port int) error {
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))
	server := &http.Server{
		Addr: addr,
		Handler: promhttp.InstrumentMetricHandler(
			r, promhttp.HandlerFor(r, promhttp.HandlerOpts{}),
		),
	}
	return httputil.ListenAndServeContext(ctx, server)
}
