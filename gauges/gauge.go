package gauges

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/prometheus/client_golang/prometheus"
)

func newGauge(
	db *sql.DB,
	opts prometheus.GaugeOpts,
	query string,
	params ...string,
) prometheus.GaugeFunc {
	iparams := make([]interface{}, len(params))
	for i, v := range params {
		iparams[i] = v
	}
	return prometheus.NewGaugeFunc(
		opts,
		func() (result float64) {
			ctx, cancel := context.WithDeadline(
				context.Background(),
				time.Now().Add(1*time.Second),
			)
			defer func() {
				<-ctx.Done()
			}()
			if err := db.QueryRowContext(ctx, query, iparams...).Scan(&result); err != nil {
				log.WithError(err).Warnf("%s: failed to query metric", opts.Name)
			}
			cancel()
			return
		},
	)
}

func isPG96(version string) bool {
	return strings.HasPrefix(version, "9.6.")
}