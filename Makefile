build:
	go build -o ./bin/gorestfaker ./cmd/gorestfaker

build_for_docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/gorestfaker-docker -a -tags netgo -ldflags '-w' ./cmd/gorestfaker

docker: build_for_docker
	docker build . -t "polarn/gorestfaker:0.0.1"
	docker tag "polarn/gorestfaker:0.0.1" "polarn/gorestfaker:latest"
