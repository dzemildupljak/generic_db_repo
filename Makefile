run-test:
	go test -coverpkg=./... ./... -v


run-test-coverage:
	go test -coverprofile=./test-coverage/cover.out ./...
	go tool cover -html=./test-coverage/cover.out -o ./test-coverage/html/cover.html