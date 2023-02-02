APP=m800-line-bot
dlv=--listen=:2345 --headless=true --api-version=2 --accept-multiclient

docker-up:
	docker compose  up -d ;

docker-down:
	docker compose down;

dlv:
	 go build -o $(APP) -gcflags "all=-N -l"
	dlv $(dlv) exec ./$(APP) m800LineBot