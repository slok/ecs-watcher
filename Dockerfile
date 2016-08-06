FROM golang:1.7-alpine


RUN apk --update add tar git bash wget && rm -rf /var/cache/apk/*

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup -g $gid ecs-watcher
RUN adduser -D -u $uid -G ecs-watcher ecs-watcher

RUN mkdir -p /go/src/github.com/slok/ecs-watcher/
RUN chown -R ecs-watcher:ecs-watcher /go

WORKDIR /go/src/github.com/slok/ecs-watcher/

USER ecs-watcher

# Install dependency manager
RUN go get github.com/Masterminds/glide
