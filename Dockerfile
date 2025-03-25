FROM golang:alpine

WORKDIR /app

ARG VERSION=0.0.0


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o persistence .

FROM bluenviron/mediamtx:latest-ffmpeg

WORKDIR /app

COPY start ./start
RUN chmod +x ./start

COPY --from=0 /app/persistence ./
RUN chmod +x ./persistence


EXPOSE 8080

CMD ["./start"]
