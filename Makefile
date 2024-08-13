build-linux:
	GOOS=linux GOARCH=amd64 go build -o togo cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o togo.exe cmd/main.go