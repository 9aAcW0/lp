FROM golang:latest

ENV GOPATH /go

RUN apt-get update && \
    apt-get upgrade -y

# install revel and revel-cli
RUN go get github.com/gin-gonic/gin

RUN ls bin

RUN mkdir /go/src/openRPA-pdf-module-go
RUN mkdir /go/src/openRPA-basic-module
WORKDIR /go/src/openRPA-pdf-module-go
WORKDIR /go/src/openRPA-basic-module
COPY ./ /go/src/openRPA-pdf-module-go
COPY ./ /go/src/openRPA-basic-module

RUN go get -v ./...

CMD ["go", "run", "app/main.go"]

EXPOSE 8080