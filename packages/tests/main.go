package tests

import (
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/prometheus-community/pro-bing"
	"math"
	"time"
	"log"
)

func Latency(host string) (int, float64, float64) {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		log.Println(err)
	}
	
	pinger.Count = 10
	pinger.SetPrivileged(true)
	pinger.Timeout = (time.Second * 10)

	roud_trip_times := []time.Duration{}
	pinger.OnRecv = func(pkt *probing.Packet) {
		roud_trip_times = append(roud_trip_times, pkt.Rtt)
	}

	lost_packets := 0
	pinger.OnFinish = func(stats *probing.Statistics) {
		lost_packets = stats.PacketsSent - stats.PacketsRecv
	}

	err = pinger.Run()
	if err != nil {
		log.Println(err)
	}

	latency_avarage := float64(0)
	jitter_avarage := float64(0)
	length_of_roud_trip_times := len(roud_trip_times)
	if length_of_roud_trip_times > 0 {
		for index, rtt := range roud_trip_times {
			microseconds := rtt.Microseconds()

			latency_avarage += float64(microseconds)
			if index > 0 {
				jitter_avarage += float64(microseconds - roud_trip_times[index-1].Microseconds())
			}
		}
		latency_avarage = (latency_avarage / float64(length_of_roud_trip_times)) / 1000
		jitter_avarage = (math.Abs(jitter_avarage / float64(length_of_roud_trip_times - 1))) / 1000
	}

	return lost_packets, latency_avarage, jitter_avarage
}

func Speed() (float64, float64) {
	server_list, err := speedtest.FetchServers()
	if err != nil {
		log.Println(err)
	}

	targets, err := server_list.FindServer([]int{})
	if err != nil {
		log.Println(err)
	}

	server := targets[0]
	err = server.DownloadTest()
	if err != nil {
		log.Println(err)
	}

	err = server.UploadTest()
	if err != nil {
		log.Println(err)
	}

	return server.DLSpeed.Mbps(), server.ULSpeed.Mbps()
}