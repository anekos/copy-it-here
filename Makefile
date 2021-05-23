

.PHONY: build
build:
	GOOS=windows go build .
	GOOS=darwin go build .
