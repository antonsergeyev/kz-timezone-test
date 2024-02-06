build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tz-linux main.go
	CGO_ENABLED=0 GOOS=darwin go build -o bin/tz-mac main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/tz-win main.go
