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

var influx client.Client

func init() {
	var err error
	influx, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("Error creating InfluxDB Client: ", err.Error())
	}
}

// CreateRouteHitMetric creates a data point to register a route hit
func createRouteHitMetric(routeName string) (*client.Point, error) {
	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"value": 1,
	}
	return client.NewPoint("hits", tags, fields, time.Now())
}

// CreateDurationMetric creates a data point to register execution time
func createDurationMetric(routeName string, duration time.Duration) (*client.Point, error) {
	elapsedTimeMs := duration.Seconds() * 1000

	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"duration": elapsedTimeMs,
	}
	return client.NewPoint("duration_ms", tags, fields, time.Now())
}

// CreateKeyQueryMetric creates a data point to register a key has been queried
func createKeyQueryMetric(keyName string) (*client.Point, error) {
	// Create a point
	tags := map[string]string{"keyName": keyName}
	fields := map[string]interface{}{
		"value": 1,
	}
	return client.NewPoint("key_query", tags, fields, time.Now())
}

// LogMetrics sends to influxdb all data points passed as parameters
func writeMetrics(points []*client.Point) {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: db,
	})

	for _, point := range points {
		bp.AddPoint(point)
	}

	// Write the batch
	influx.Write(bp)
}

// LogMetrics log metrics to the default data store
func LogMetrics(routeName, key string, duration time.Duration) {
	points := []*client.Point{}
	if hitPoint, err := createRouteHitMetric(routeName); err == nil {
		points = append(points, hitPoint)
	}
	if durationPoint, err := createDurationMetric(routeName, duration); err == nil {
		points = append(points, durationPoint)
	}
	if len(key) > 0 {
		if keyQueryPoint, err := createKeyQueryMetric(key); err == nil {
			points = append(points, keyQueryPoint)
		}
	}

	writeMetrics(points)
}
