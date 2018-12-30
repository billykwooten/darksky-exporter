# Create Builder
FROM golang:1.11.4-alpine3.8 as builder

RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser

# Create Container
FROM scratch

MAINTAINER Billy Wooten <billykwooten@gmail.com>

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
ADD darksky-exporter /app

USER appuser
ENTRYPOINT ["/app"]
