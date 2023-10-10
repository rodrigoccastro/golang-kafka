In this repository we have a mock of 3 microservices 
In the first microservice we user kafka


1. run service kafka in docker 
	in path ./
	docker-compose up -d

2. run producer = send to kafka
	in path ./producer execute: go run .

3. run consumer = read from kafka
	in path ./consumer execute: go run .
