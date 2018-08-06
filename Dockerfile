FROM golang:1.8

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build -o /app/main .
CMD ["/app/main"]
