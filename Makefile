run:
	./swag init
	go run .

container:
	docker-compose build
	docker-compose up

test:
	go test *.go

build:
	go build . -o server