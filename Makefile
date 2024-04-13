# Copyright (c) 2022 Isaque Veras
# Use of this source code is governed by MIT style
# license that can be found in the LICENSE file.

.PHONY: build run local down-local docker-clean logs-local

FILES := $(shell docker ps -aq)

# - trimpath 	- will remove the filepathes from the reports, good to same money on network trafic,
#             	focus on bug reports, and find issues fast.
# - race    	- adds a racedetector, in case of racecondition, you can catch report with sentry.
#             	https://golang.org/doc/articles/race_detector.html
build: ## Builds binary
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o powersso \
		./
	@ echo "done"

.ONESHELL:
image-build: ## Docker Build
	@ echo "Docker Build"
	@ DOCKER_BUILDKIT=0 docker build \
		--file Dockerfile \
		--tag powersso \
			.

run:
	@ go run main.go

test:
	@ go test ./...

dev:
	@ docker compose -f ./docker-compose.dev.yml up -d 

local:
	@ docker compose -f ./docker-compose.local.yml up -d 

local-build:
		@ docker compose -f ./docker-compose.local.yml up -d --build

down-local:
	@ docker stop $(FILES)
	@ docker rm $(FILES)

docker-clean:
	@ docker system prune -f

logs-local:
	@ docker logs -f $(FILES)

lint:
	golangci-lint run ./...

check:
	staticcheck ./...
