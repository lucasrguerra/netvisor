from influxdb_client.client.write_api import SYNCHRONOUS
from influxdb_client import InfluxDBClient
import os

class Database:
    def __init__(self):
        influx_url = os.getenv("INFLUXDB_URL")
        influx_token = os.getenv("INFLUXDB_TOKEN")
        self.org = os.getenv("INFLUXDB_ORG")
        self.bucket = os.getenv("INFLUXDB_BUCKET")
        self.client = InfluxDBClient(url=influx_url, token=influx_token, org=self.org)

    def useAPI(self, data):
        write_api = self.client.write_api(write_options=SYNCHRONOUS)
        write_api.write(self.bucket, self.org, data)
        write_api.close()

    def insertLatencyTest(self, host, latency, jitter, failed_pings):
        self.useAPI([
            {
                "measurement": "latency",
                "tags": {
                    "host": host
                },
                "fields": {
                    "latency": latency,
                    "jitter": jitter,
                    "failed_pings": failed_pings
                }
            }
        ])

    def insertSpeedTest(self, download_speed, upload_speed):
        self.useAPI([
            {
                "measurement": "speedtest",
                "fields": {
                    "download_speed": download_speed,
                    "upload_speed": upload_speed
                }
            }
        ])