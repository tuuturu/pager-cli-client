.PHONY: build/pager-client

build/pager-client:
	mkdir -p ./build
	go build -o ./build/pager-client main.go

build: build/pager-client

install:
	mkdir -p ~/.local/bin
	cp ./build/pager-client ~/.local/bin/pager-client

docker-image:
	docker build \
		--tag docker.pkg.github.com/tuuturu/pager-cli-client/pager-cli-client:local \
		.

run-docker:
	docker run \
		-e DISCOVERY_URL= \
		-e CLIENT_ID= \
		-e CLIENT_SECRET= \
		-e EVENTS_SERVICE_URL= \
		docker.pkg.github.com/tuuturu/pager-cli-client/pager-cli-client:local \
		"Package delivery: delivered" "Your package from Amazon has arrived"
