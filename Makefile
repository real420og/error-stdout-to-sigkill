build_linux:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o errorToSigkill_linux .
build_darwin:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s" -a -installsuffix cgo -o errorToSigkill_darwin .

test:
	go test -timeout=5s ./...