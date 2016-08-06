FROM slok/ecs-watcher_base:latest
USER root
RUN apk add --no-cache g++

USER ecs-watcher
RUN go get github.com/golang/mock/mockgen
