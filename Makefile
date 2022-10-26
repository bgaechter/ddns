run: 
	go run cmd/ddns/main.go

build: 
	go build ./... 

test: 
	go vet -v &&\
	go test -v
