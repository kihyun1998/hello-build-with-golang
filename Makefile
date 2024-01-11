build:
	go build -o $(TAG_NAME)_wadd ./wadd
	go build -o $(TAG_NAME)_wsub ./wsub

test:
	go test -v -cover ./wadd/...
	go test -v -cover ./wsub/...

.PHONY: build, test