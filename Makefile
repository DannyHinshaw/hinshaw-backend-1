run:
	docker-compose up --build -d

stop:
	docker-compose down -v --remove-orphans
