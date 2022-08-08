start:
	docker-compose up -d --remove-orphans
stop:
	docker-compose down	
producer_run:
	go run producer/main.go
consumer-1:
	go run consumer_1/main.go
consumer-2:
	go run consumer_2/main.go
consumer-3:
	go run consumer_3/main.go