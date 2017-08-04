package gauges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackends(t *testing.T) {
	var assert = assert.New(t)
	gauges, close := prepare(t)
	defer close()
	var metrics = evaluate(t, gauges.Backends())
	assert.Len(metrics, 1)
	assertGreaterThan(t, 0, metrics[0])
}

func TestMaxBackends(t *testing.T) {
	var assert = assert.New(t)
	gauges, close := prepare(t)
	defer close()
	var metrics = evaluate(t, gauges.MaxBackends())
	assert.Len(metrics, 1)
	assertGreaterThan(t, 0, metrics[0])
}

func TestWaitingBackends(t *testing.T) {
	var assert = assert.New(t)
	gauges, close := prepare(t)
	defer close()
	var metrics = evaluate(t, gauges.WaitingBackends())
	assert.Len(metrics, 1)
	assertGreaterThan(t, -1, metrics[0])
}

func TestBackendsStatus(t *testing.T) {
	var assert = assert.New(t)
	gauges, close := prepare(t)
	defer close()
	var metrics = evaluate(t, gauges.BackendsStatus()...)
	assert.Len(metrics, 3)
	for _, m := range metrics {
		assertGreaterThan(t, -1, m)
	}
}