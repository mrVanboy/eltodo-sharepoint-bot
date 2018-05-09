FROM golang:1.10 AS build-env
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /bin/dep
ADD ./src /go/src
RUN chmod +x /bin/dep && \
    cd /go/src/sharepoint-bot && \
    dep ensure -vendor-only && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goapp

# final stage
FROM alpine
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/src/sharepoint-bot/goapp /app
ENTRYPOINT ["./goapp"]
