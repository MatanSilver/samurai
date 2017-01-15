all:
	go build
install: all
	sudo cp samurai /usr/local/bin/samurai
uninstall:
	sudo rm /usr/local/bin/samurai
