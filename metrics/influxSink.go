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

func LogHit(routeName string) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: db,
	})

	// Create a point and add to batch
	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"value": 1,
	}
	pt, _ := client.NewPoint("hits", tags, fields, time.Now())
	bp.AddPoint(pt)

	// Write the batch
	influx.Write(bp)
}

func LogDuration(routeName string, duration time.Duration) {
	elapsedTimeMs := duration.Seconds() * 1000
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database: db,
	})

	// Create a point and add to batch
	tags := map[string]string{"route": routeName}
	fields := map[string]interface{}{
		"duration": elapsedTimeMs,
	}
	pt, _ := client.NewPoint("duration_ms", tags, fields, time.Now())
	bp.AddPoint(pt)

	// Write the batch
	influx.Write(bp)
}
