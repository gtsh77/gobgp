#GOBGPD
FROM golang:1.17

RUN mkdir -p /go/src
COPY src/ /go/src/gobgp/
WORKDIR /go/src/gobgp/pkg/cmd/gobgpd
RUN go build -o gobgpd

RUN mkdir /gobgpd
RUN mv gobgpd /gobgpd/.
RUN rm -rf /go/src/gobgp
WORKDIR /gobgpd