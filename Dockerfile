
#
# Build project in separate container
#

FROM golang:alpine3.13 AS build


RUN apk update && \
    apk upgrade && \
    apk add git
COPY . /go/src/github.com/slysmoke/jumpgate

WORKDIR /go/src/github.com/slysmoke/jumpgate
RUN go mod init github.com/slysmoke/jumpgate
RUN go get go.opentelemetry.io/otel/label
RUN go get -d -v ./...
RUN go build
RUN cp /go/src/github.com/slysmoke/jumpgate/jumpgate /jumpgate



FROM alpine:3.7

#
# Copy release to container and set command
#

# Add faster mirror and upgrade packages in base image, load ca certs, otherwise no TLS for us
RUN printf "http://mirror.leaseweb.com/alpine/v3.7/main\nhttp://mirror.leaseweb.com/alpine/v3.7/community" > etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

# Copy build
COPY jumpgate jumpgate

ENV LISTEN_HOST ":8000"
EXPOSE 8000

CMD ["/jumpgate"]