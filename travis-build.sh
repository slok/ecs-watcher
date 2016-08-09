#!/bin/bash -e

# Run vet
make vet

# Run tests
make test


# On successful build
if [ "$TRAVIS_BRANCH" == "master" ]; then
    make push
fi
