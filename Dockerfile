# Builder image
FROM golang:latest as builder

# Install dependencies
RUN apt-get update && apt-get upgrade -q -y && \
	apt-get install -y --no-install-recommends git build-essential

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - &&  \
	apt-get install -y nodejs 

RUN go get -u github.com/gobuffalo/packr/... && go get -u github.com/gopherjs/gopherjs && npm install -g sass

ADD . /go/src/github.com/gbotrel/zsldemo

# go generate will generate app.css from app.scss and compile app.js (gopherjs)
RUN cd /go/src/github.com/gbotrel/zsldemo/frontend && go generate 

# packr produces a fat binary including the static assets (html, css, js)
RUN cd /go/src/github.com/gbotrel/zsldemo && packr && go build

# Binaries only image
FROM debian:latest

COPY --from=builder /go/src/github.com/gbotrel/zsldemo/zsldemo /usr/local/bin/

EXPOSE 8001
ENTRYPOINT ["zsldemo"]
