build:
	go build -o auto-restart .
install:build
	cp auto-restart  /opt/homebrew/bin/auto-restart
