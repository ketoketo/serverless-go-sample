.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./bin/hello

# どちらでもstatic link buildになる
# GOBUILD=env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -o
GOBUILD=env GOOS=linux CGO_ENABLED=0 go build
build:
	$(GOBUILD) -o bin/hello hello-world/main.go
	$(GOBUILD) -o bin/clients functions/clients/main.go

debug:
	$(GOBUILD) -o dlv github.com/derekparker/delve/cmd/dlv
	$(GOBUILD) -gcflags='-N -l' -o hello-world/hello-world ./hello-world
	sam local start-api -d 5986 --debugger-path . --debug-args="-delveAPI=2"

test:
	go test ./...

deploy:
	sls deploy -v

remove:
	sls remove -v