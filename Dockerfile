# Create Builder & Build
FROM golang:1.11.4-alpine3.8 as builder

RUN apk update && apk add --no-cache git && apk add ca-certificates
RUN adduser -D -g '' appuser

COPY . $GOPATH/src/github.com/billykwooten/darksky-exporter
WORKDIR $GOPATH/src/github.com/billykwooten/darksky-exporter

RUN go get golang.org/x/tools/cmd/cover
RUN	go get github.com/go-playground/overalls
RUN	overalls -covermode=atomic -project=github.com/billykwooten/darksky-exporter -- -race -v
RUN mv overalls.coverprofile darksky-exporter.cover
RUN	go tool cover -func=darksky-exporter.cover

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -a -installsuffix cgo -ldflags="-w -s" -o /app

# Create Container
FROM scratch

MAINTAINER Billy Wooten <billykwooten@gmail.com>

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app /app

USER appuser
ENTRYPOINT ["/app"]
