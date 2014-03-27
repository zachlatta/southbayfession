FROM ubuntu:13.10

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install golang git mercurial build-essential
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN go get github.com/codegangsta/gin
RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN go get github.com/zachlatta/southbayfession
WORKDIR /go/src/github.com/zachlatta/southbayfession
ADD . /go/src/github.com/zachlatta/southbayfession
RUN go get
