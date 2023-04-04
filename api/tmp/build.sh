#!/bin/bash

# daocker版本
# docker build -t flashexpress/bi_web  .
# #docker-compose down
# docker-compose up -d
# #多层编译会产生一个 none 镜像
# docker images | grep "none" |awk '{print $3}' |xargs docker rmi 

# 原生版本
cd ../
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o name.bin main.go
rsync -r --progress name.bin wisdom:/www/wwwroot/name-api.wisdom-os.top
# ./build.sh && scp i18n.bin 110:/tmp  && ssh 110 " cd /mnt/www/i18n-api && mv i18n.bin i18n.bin.bk && mv /tmp/i18n.bin ."