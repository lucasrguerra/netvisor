from sources import database
from sources import latency
from sources import utils
from sources import speed
import time

def main():
    utils.loadEnvirontmentVariables()

    db = database.Database()

    host_list = [
        "1.1.1.1",
        "8.8.8.8",
        "10.10.10.10",
        "a.dns.br",
    ]

    while True:
        for host in host_list:
            latency_average, jitter_average, failed = latency.makeLatencyTest(host, 10)
            db.insertLatencyTest(host, latency_average, jitter_average, failed)
            print("Host: ", host)

        download_speed, upload_speed = speed.makeSpeedTest()
        print("Download Speed: ", download_speed)
        print("Upload Speed: ", upload_speed)
        db.insertSpeedTest(download_speed, upload_speed)
        
        time.sleep(600)



if __name__ == '__main__':
    main()
