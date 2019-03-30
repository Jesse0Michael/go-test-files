
build:
	if [ ! -d cmd ]; then mkdir cmd; fi
	go build -o cmd/testfiles
