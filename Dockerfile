FROM golang:1.11.4

ARG LD_FLAGS=""

RUN go version

RUN go get github.com/mitchellh/gox
RUN go get github.com/go-playground/overalls

ENV GOAPP_PATH /go/src/github.com/billykwooten/darksky-exporter

RUN mkdir -p $GOAPP_PATH
ADD . $GOAPP_PATH
WORKDIR $GOAPP_PATH

RUN ./script/test

RUN go build -ldflags "$LD_FLAGS" .

RUN gox -ldflags "$LD_FLAGS" -os="darwin linux windows" -arch="amd64" -output "darksky-exporter-{{.OS}}-{{.Arch}}"

CMD tar -cf - darksky-exporter*
