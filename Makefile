.PHONY: docs
docs:
	redoc-cli build docs/index.yaml -o docs/index.html
build:
	make docs
	mkdir -p tmp
	GOOS=linux GOARCH=arm64 go build -o tmp/bootstrap 
	cd tmp && zip -j bootstrap.zip bootstrap && cd ..
