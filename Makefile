SHELL := /bin/bash
.PHONY: deps create install-linux-global install-linux install-mac

create:
	cd ./utdf && \
	GOOS=windows go build -o ../release/windows/utdf.exe && \
	GOOS=linux go build -o ../release/linux/utdf && \
	go build -o ../release/mac/utdf

deps:
	mkdir -p ~/bin  && \
	./release/scripts/setup.sh 
	source ~/.bashrc

install-linux-global:
	cd ./release && \
	cp ./linux/utdf /usr/local/bin

install-linux: deps
	cd ./release && \
	cp ./linux/utdf ~/bin
        

install-mac:
	cd ./release && \
	cp ./mac/utdf /usr/local/bin