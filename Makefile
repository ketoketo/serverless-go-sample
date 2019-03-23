.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./bin/hello

# どちらでもstatic link buildになる
# GOBUILD=env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -o
GOBUILD=env GOOS=linux CGO_ENABLED=0 go build -o
build:
	$(GOBUILD) bin/hello hello-world/main.go

test:
	go test ./...

deploy:
	sls deploy -v

remove:
	sls remove -v