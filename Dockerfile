FROM golang:latest
WORKDIR $GOPATH/src/github.com/tonyjhang/goRestfulPractice
COPY . $GOPATH/src/github.com/tonyjhang/goRestfulPractice
RUN go build .
EXPOSE 8888
ENTRYPOINT ["./goRestfulPractice"]