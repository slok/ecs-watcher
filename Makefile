# The following are targers that do not exist in the filesystem as real files and should be always executed by make
.PHONY: default deps login base build dev shell start stop image test

# Name of this service/application
SERVICE_NAME := ecs-watcher

# Docker image name for this project
IMAGE_NAME := slok/$(SERVICE_NAME)

# Shell to use for running scripts
SHELL := $(shell which bash)

# Get docker path or an empty string
DOCKER := $(shell command -v docker)

# Get docker-compose path or an empty string
DOCKER_COMPOSE := $(shell command -v docker-compose)

# Get the main unix group for the user running make (to be used by docker-compose later)
GID := $(shell id -g)

# Get the unix user id for the user running make (to be used by docker-compose later)
UID := $(shell id -u)

# Get the username of theuser running make. On the devbox, we give priority to /etc/username
USERNAME ?= $(shell ( [ -f /etc/username ] && cat /etc/username  ) || whoami)

# File to keep track of the last login to the docker registry, so that login is not ran every time
LOGIN_FILE := ~/.devlogin

# Bash history file for container shell
HISTORY_FILE := ~/.bash_history.$(SERVICE_NAME)

# Get git tag
TAG := $(shell git describe --tags --exact-match)
ifndef TAG
	TAG := latest
endif

# Remove login flag if older than 10 hours
_ := $(shell find $(LOGIN_FILE) -mmin +600 -delete)

# The default action of this Makefile is to build the development docker image
default: build

# Test if the dependencies we need to run this Makefile are installed
deps:
ifndef DOCKER
	@echo "Docker is not available. Please install docker"
	@exit 1
endif
ifndef DOCKER_COMPOSE
	@echo "docker-compose is not available. Please install docker-compose"
	@exit 1
endif

# Build the login file, previously doing a proper login
$(LOGIN_FILE):
	touch $(LOGIN_FILE)

# Alias just for having a better ui
login: $(LOGIN_FILE)

# Build the base docker image which is shared between the development and production images
base: deps login
	docker build -t $(IMAGE_NAME)_base:latest .

# Build the development docker image
build: base
	cd environment/dev && docker-compose build

# Run the development environment in non-daemonized mode (foreground)
dev: build
	cd environment/dev && \
	( docker-compose up; \
		docker-compose stop; \
		docker-compose rm -f; )

# Run a shell into the development docker image
shell: build
	-touch $(HISTORY_FILE)
	cd environment/dev && docker-compose run --service-ports --rm $(SERVICE_NAME) /bin/bash

# Run the development environment in the background
start: build
	cd environment/dev && \
		docker-compose up -d

# Stop the development environment (background and/or foreground)
stop:
	cd environment/dev && ( \
		docker-compose stop; \
		docker-compose rm -f; \
		)

# Build release, target on /bin
build_release:build
		cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c "go build -o ./bin/ecs-watcher --ldflags '-w -linkmode external -extldflags \"-static\"'  ./ "

# Update project dependencies to tle latest version
dep_update:build
	cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c 'glide up --strip-vcs --update-vendored'
# Install new dependency make dep_install args="github.com/Sirupsen/logrus"
dep_install:build
		cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c 'glide get --strip-vcs $(args)'

# Pass the golang vet check
vet: build
	cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c 'go vet `glide nv`'

# Execute unit tests
test:build
	cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c 'go test `glide nv` -v'

# Generate required code (mocks...)
gogen: build
	cd environment/dev && docker-compose run --rm $(SERVICE_NAME) /bin/bash -c 'go generate `glide nv`'

# Build the production image
image:
	docker build \
	-t $(IMAGE_NAME) \
	-t $(IMAGE_NAME):$(TAG) \
	-f environment/prod/Dockerfile \
	./

# Push the production docker image to the repository
#push: image
#	docker push $(REPOSITORY)
#
