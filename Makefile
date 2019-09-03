.PHONY: lint
lint:
	golangci-lint run . 
.PHONY: test
test:
	go test -test.count=1 -test.v=true ./...
.PHONY: mute-test
mute-test:
	go-mutesting --do-not-remove-tmp-folder
