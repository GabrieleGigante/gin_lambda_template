mkdir -p tmp
GOOS=linux GOARCH=arm64 go build -o tmp/bootstrap 
./generate-docs.sh
cd tmp && zip -j bootstrap.zip bootstrap && cd ..