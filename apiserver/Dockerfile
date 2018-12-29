FROM golang:latest

WORKDIR $GOPATH/src/github.com/haibeichina/gin_video_server/apiserver
COPY . $GOPATH/src/github.com/haibeichina/gin_video_server/apiserver
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./apiserver"]