HOME_DIR = $(shell cd ~ && pwd)
PATH = $$PATH:$(HOME_DIR)/bin
.PHONY: create

create:
	cd ./utdf && \
	GOOS=windows go build -o ../release/windows/utdf.exe && \
	GOOS=linux go build -o ../release/linux/utdf && \
	go build -o ../release/mac/utdf

deps:
	mkdir -p $(HOME_DIR)/bin
	export PATH

install-linux-global:
	cd ./release && \
	cp ./linux/utdf /usr/local/bin

install-linux: deps
	cd ./release && \
	cp ./linux/utdf $(HOME_DIR)/bin

install-mac:
	cd ./release && \
	cp ./mac/utdf /usr/local/bin
