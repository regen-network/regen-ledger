package keeper

import "github.com/prometheus/client_golang/prometheus"

func WithVMCacheMetrics(prometheus.Registerer) Option {
	return Option{}
}
