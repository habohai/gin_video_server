FROM golang:latest

WORKDIR $GOPATH/src/github.com/haibeichina/gin_video_server/web
COPY . $GOPATH/src/github.com/haibeichina/gin_video_server/web
RUN go build .

EXPOSE 8085
ENTRYPOINT ["./web"]