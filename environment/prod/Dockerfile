FROM slok/ecs-watcher_base:latest

USER root

# Prepare
WORKDIR /go/src/github.com/slok/ecs-watcher/
COPY . ./
RUN mkdir -p /bin

# Build
RUN go build -o /bin/ecs-watcher ./
RUN chmod 755 /bin/ecs-watcher

# Clean up
WORKDIR /
RUN rm -rf /go/src/*


ENTRYPOINT [ "/bin/ecs-watcher" ]
