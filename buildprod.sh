#! /bin/bash

# use for local test

rm -rf ~/goprojects/bin/runtime/templates
mkdir ~/goprojects/bin/runtime/templates
cp ~/goprojects/src/github.com/haibeichina/gin_video_server/streamserver/runtime/templates/upload.html ~/goprojects/bin/runtime/templates
cp -R ~/goprojects/src/github.com/haibeichina/gin_video_server/web/runtime/templates/* ~/goprojects/bin/runtime/templates


rm ~/goprojects/bin/conf/apiserver.ini
rm ~/goprojects/bin/apiserver
cp ~/goprojects/src/github.com/haibeichina/gin_video_server/apiserver/conf/apiserver.ini ~/goprojects/bin/conf/
cd ~/goprojects/src/github.com/haibeichina/gin_video_server/apiserver
go build -o ~/goprojects/bin/apiserver

rm ~/goprojects/bin/conf/scheduler.ini
rm ~/goprojects/bin/scheduler
cp ~/goprojects/src/github.com/haibeichina/gin_video_server/scheduler/conf/scheduler.ini ~/goprojects/bin/conf/
cd ~/goprojects/src/github.com/haibeichina/gin_video_server/scheduler
go build -o ~/goprojects/bin/scheduler

rm ~/goprojects/bin/conf/streamserver.ini
rm ~/goprojects/bin/streamserver
cp ~/goprojects/src/github.com/haibeichina/gin_video_server/streamserver/conf/streamserver.ini ~/goprojects/bin/conf/
cd ~/goprojects/src/github.com/haibeichina/gin_video_server/streamserver
go build -o ~/goprojects/bin/streamserver

rm ~/goprojects/bin/conf/web.ini
rm ~/goprojects/bin/web
cp ~/goprojects/src/github.com/haibeichina/gin_video_server/web/conf/web.ini ~/goprojects/bin/conf/
cd ~/goprojects/src/github.com/haibeichina/gin_video_server/web
go build -o ~/goprojects/bin/web