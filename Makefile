
build:
	if [ ! -d bin ]; then mkdir bin; fi
	go build -o bin/gofiles github.com/jesse0michael/go-test-files/gofiles
