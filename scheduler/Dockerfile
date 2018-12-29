FROM golang:latest

WORKDIR $GOPATH/src/github.com/haibeichina/gin_video_server/scheduler
COPY . $GOPATH/src/github.com/haibeichina/gin_video_server/scheduler
RUN go build .

EXPOSE 9001
ENTRYPOINT ["./scheduler"]