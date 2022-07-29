run:
	go run main.go

docker:
	docker-compose up 

docker-build:
	docker-compose up --build
	
.PHONY: run docker docker-build