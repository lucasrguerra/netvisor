package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
	"os"
)

var client influxdb2.Client

func Init() {
	influxdb_url := os.Getenv("INFLUXDB_URL")
	influxdb_token := os.Getenv("INFLUXDB_TOKEN")

	client = influxdb2.NewClient(influxdb_url, influxdb_token)
}

func Close() {
	client.Close()
}

func InsertLatencyTest(host string, latency float64, jitter float64, failed_pings int) {
	influxdb_organization := os.Getenv("INFLUXDB_ORGANIZATION")
	influxdb_bucket := os.Getenv("INFLUXDB_BUCKET")
	write_api := client.WriteAPI(influxdb_organization, influxdb_bucket)

	point := influxdb2.NewPointWithMeasurement("latency").
		AddTag("host", host).
		AddField("latency", latency).
		AddField("jitter", jitter).
		AddField("failed_pings", failed_pings).
		SetTime(time.Now())

	write_api.WritePoint(point)
	write_api.Flush()
}

func InsertSpeedTest(download_speed float64, upload_speed float64) {
	influxdb_organization := os.Getenv("INFLUXDB_ORGANIZATION")
	influxdb_bucket := os.Getenv("INFLUXDB_BUCKET")
	write_api := client.WriteAPI(influxdb_organization, influxdb_bucket)

	point := influxdb2.NewPointWithMeasurement("speed").
		AddField("download_speed", download_speed).
		AddField("upload_speed", upload_speed).
		SetTime(time.Now())

	write_api.WritePoint(point)
	write_api.Flush()
}