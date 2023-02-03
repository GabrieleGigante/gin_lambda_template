mkdir -p tmp
GOOS=linux GOARCH=arm64 go build -o tmp/bootstrap main.go
zip -j tmp/bootstrap.zip tmp/bootstrap