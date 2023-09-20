help:
	@echo "Available commands: run, stop"

run:
	go build -C ./src/consumer -o ../../build/consumer/compiled && \
	go build -C ./src/producer -o ../../build/producer/compiled

	docker compose up -d --build
	
stop: 
	docker compose down
