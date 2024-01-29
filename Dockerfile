#--- Build stage
FROM golang:1.20-bullseye AS go-builder

WORKDIR /src

COPY . /src/

RUN make build CGO_ENABLED=0

#--- Image stage
FROM alpine:3.19.1

COPY --from=go-builder /src/target/dist/nemeton-leaderboard /usr/bin/nemeton-leaderboard

WORKDIR /opt

ENTRYPOINT ["/usr/bin/nemeton-leaderboard"]
