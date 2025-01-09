package main

import (
    "netvisor/packages/database"
	"netvisor/packages/tests"
	"netvisor/packages/utils"
    "strconv"
	"time"
	"log"
    "os"
)

func main() {
    log.Println("NetVisor is running!")
    utils.LoadEnvironmentVariables()
    database.Init()

    seconds_between_tests, _ := strconv.Atoi(os.Getenv("SECONDS_BETWEEN_TESTS"))
    for {
        hosts_list := []string{
            "1.1.1.1",
            "8.8.8.8",
            "10.10.10.10",
            "a.dns.br",
        }

        for _, host := range hosts_list {
            lost_packets, latency_avarage, jitter_avarage := tests.Latency(host)
            log.Printf("Host: %s / Lost packets: %d / Latency avarage: %.2f ms / Jitter avarage: %.2f ms\n", host, lost_packets, latency_avarage, jitter_avarage)
            database.InsertLatencyTest(host, latency_avarage, jitter_avarage, lost_packets)
        }

        download_speed, upload_speed := tests.Speed()
        log.Printf("Download speed: %.2f Mbps / Upload speed: %.2f Mbps\n", download_speed, upload_speed)
        database.InsertSpeedTest(download_speed, upload_speed)

        time.Sleep(time.Second * time.Duration(seconds_between_tests))
    }
}