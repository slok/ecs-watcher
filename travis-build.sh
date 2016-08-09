#!/bin/bash -e

# Run vet
make vet

# Run tests
make test


# On successful build
if [ "$TRAVIS_BRANCH" == "master" ]; then
    docker login -e $DOCKER_EMAIL -u $DOCKER_LOGIN -p $DOCKER_PASSWORD
    docker push slok/ecs-watcher;
fi
