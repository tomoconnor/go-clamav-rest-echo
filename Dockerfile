FROM golang:1.18 as builder

WORKDIR /go/src/github.com/tomoconnor/go-clamav-rest-echo

COPY go.mod go.sum main.go server.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-clamav-rest-echo

FROM alpine:3.16
RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -S app -G app

USER 1000

COPY --from=builder /go/src/github.com/tomoconnor/go-clamav-rest-echo/go-clamav-rest-echo /go-clamav-rest-echo

EXPOSE 8080
ENV CLAMD_HOST localhost
ENV CLAMD_PORT 3310
ENV LISTEN_PORT 8080

ENTRYPOINT ["/go-clamav-rest-echo"]
