package metrics

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/rcrowley/go-metrics"
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/vrischmann/go-metrics-influxdb"
	"log"
	"time"
)

var GetValueCounter = metrics.NewCounter()
var AddValueCounter = metrics.NewCounter()
var GetValueRequestsTimer = metrics.NewTimer()
var AddValueRequestsTimer = metrics.NewTimer()
var RequestsMeter = metrics.NewMeter()
var nodesRemovedCounter = metrics.NewCounter()

func init() {
	metrics.Register("getValue_endpoint_counter", GetValueCounter)
	metrics.Register("addValue_endpoint_counter", AddValueCounter)
	metrics.Register("getValue_requests_timer", GetValueRequestsTimer)
	metrics.Register("addValue_requests_timer", AddValueRequestsTimer)
	metrics.Register("requests_rate", RequestsMeter)
	metrics.Register("deleted_nodes_counter", nodesRemovedCounter)
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go influxdb.InfluxDB(
		metrics.DefaultRegistry, // metrics registry
		time.Second*10,          // interval
		"http://localhost:8086", // the InfluxDB url
		"goCacheIt",             // your InfluxDB database
		"goCacheItUser",         // your InfluxDB user
		"test",                  // your InfluxDB password
	)
}

func LogHit(endpointName string) {
	RequestsMeter.Mark(1)
	switch endpointName {
	case "getValue":
		GetValueCounter.Inc(1)
		log.Printf("Hit endpoint %s. Request count %d\n", endpointName,
			GetValueCounter.Count())
	case "addValue":
		AddValueCounter.Inc(1)
		log.Printf("Hit endpoint %s. Request count %d\n", endpointName,
			AddValueCounter.Count())
	}
	log.Printf("Average req/s: %f ", RequestsMeter.RateMean())
}

func LogNodesRemoval(count int) {
	nodesRemovedCounter.Inc(int64(count))
	log.Printf("Removed %d nodes.\n", count)
}

func ResetNodesRemovalCounter() {
	nodesRemovedCounter.Clear()
}
