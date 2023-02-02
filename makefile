build:
	make test
	docker build -t tcgraf .

run:
	make build
	docker run -d -p 9999:80 tcgraf

test:
	go test -cover ./...
