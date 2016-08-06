#!/bin/bash -e

# Run vet
make vet

# Run tests
make test


# On successful build, upload the image if this is not a PR and this is a master build
#if [ "${TRAVIS_PULL_REQUEST}" == "false" ]; then
#    BRANCH=${TRAVIS_BRANCH} make push
#fi
