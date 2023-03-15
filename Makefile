docker-build:
	docker build -t countries-telegram-bot .
docker-run:
	docker run -p 8080:8080 -it countries-telegram-bot
