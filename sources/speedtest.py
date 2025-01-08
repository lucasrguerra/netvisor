import speedtest

def makeSpeedTest():
    st = speedtest.Speedtest()
    st.get_best_server()
    download_speed = st.download() / 1_000_000 
    upload_speed = st.upload() / 1_000_000
    return download_speed, upload_speed