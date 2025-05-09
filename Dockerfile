FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o persistence .

FROM bluenviron/mediamtx:latest-ffmpeg


RUN mkdir -p /opt/startup

COPY start /start
RUN chmod +x /start

COPY --from=0 /app/persistence /opt/startup/persistence
RUN chmod +x /opt/startup/persistence

COPY mediamtx.yaml /opt/startup/mediamtx.yaml
RUN chmod +x /opt/startup/mediamtx.yaml



ENTRYPOINT [ "/start" ]
