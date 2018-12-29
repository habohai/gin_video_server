FROM golang:latest

WORKDIR $GOPATH/src/github.com/haibeichina/gin_video_server/streamserver
COPY . $GOPATH/src/github.com/haibeichina/gin_video_server/streamserver
RUN go build .

EXPOSE 9000
ENTRYPOINT ["./streamserver"]