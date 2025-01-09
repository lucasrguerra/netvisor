FROM arm32v7/alpine:latest

WORKDIR /app

COPY ./netvisor /app/netvisor

RUN chmod +x /app/netvisor

CMD ["/app/netvisor"]
