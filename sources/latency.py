from ping3 import ping
import time

def makeLatencyTest(host, count):
    latencys = []
    jitter = []
    failed = 0

    for i in range(count):
        latency = ping(host, unit='ms')
        if latency is not None:
            latencys.append(latency)
        else:
            failed += 1
        time.sleep(1)

    if len(latencys) == 0:
        return None, None
    
    latency_average = sum(latencys) / len(latencys)
    for i in range(1, len(latencys)):
        jitter.append(abs(latencys[i] - latencys[i - 1]))
    
    jitter_average = sum(jitter) / len(jitter)
    return latency_average, jitter_average, failed

