package metrics

import (
	"github.com/influxdb/influxdb/client/v2"
	"log"
	"time"
)

const (
	db = "goCacheIt"
	// username = "bubba"
	// password = "bumblebeetuna"
)

// Type declaration
type metrics struct {
	points []*client.Point
}

// Global variables
var influx client.Client
var metricsChannel chan *metrics
var emptySinkTicker *time.Ticker
var points []*client.Point

func init() {
	var err error
	influx, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("Error creating InfluxDB Client: ", err.Error())
	}
	metricsChannel = make(chan *metrics)
	points = []*client.Point{}
	emptySinkTicker = time.NewTicker(10 * time.Second)
	go worker()
}

func worker() {
	for {
		select {
		case metrics, more := <-metricsChannel:
			if metrics != nil && metrics.points != nil && len(metrics.points) > 0 {
				for _, point := range metrics.points {
					points = append(points, point)
				}
			}
			if !more {
				return
			}
		case <-emptySinkTicker.C:
			sendBatchMetrics()
			points = []*client.Point{}
		}
	}
}

// sendMetrics sends to influxdb all data points passed as parameters
func sendBatchMetrics() {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: db,
	})

	for _, point := range points {
		bp.AddPoint(point)
	}

	// Write the batch
	influx.Write(bp)
}

// createRouteHitMetric creates a data point to register a route hit
func createRouteHitMetric(routeName string) (*client.Point, error) {
	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"value": 1,
	}
	return client.NewPoint("hits", tags, fields, time.Now())
}

// createDurationMetric creates a data point to register execution time
func createDurationMetric(routeName string, duration time.Duration) (*client.Point, error) {
	elapsedTimeMs := duration.Seconds() * 1000

	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"duration": elapsedTimeMs,
	}
	return client.NewPoint("duration_ms", tags, fields, time.Now())
}

// createKeyQueryMetric creates a data point to register a key has been queried
func createKeyQueryMetric(keyName string) (*client.Point, error) {
	// Create a point
	tags := map[string]string{"keyName": keyName}
	fields := map[string]interface{}{
		"value": 1,
	}
	return client.NewPoint("key_query", tags, fields, time.Now())
}

// LogMetrics logs metrics to the configured influxdb instance
func LogMetrics(routeName, key string, duration time.Duration) {
	collectedMetrics := &metrics{
		points: []*client.Point{},
	}
	if hitPoint, err := createRouteHitMetric(routeName); err == nil {
		collectedMetrics.points = append(collectedMetrics.points, hitPoint)
	}
	if durationPoint, err := createDurationMetric(routeName, duration); err == nil {
		collectedMetrics.points = append(collectedMetrics.points, durationPoint)
	}
	if len(key) > 0 {
		if keyQueryPoint, err := createKeyQueryMetric(key); err == nil {
			collectedMetrics.points = append(collectedMetrics.points, keyQueryPoint)
		}
	}
	metricsChannel <- collectedMetrics
}
