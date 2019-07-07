.PHONY: create

create:
	cd ./utdf && \
	GOOS=windows go build -o ../release/windows/utdf.exe && \
	GOOS=linux go build -o ../release/linux/utdf && \
	go build -o ../release/mac/utdf

install-linux:
	cd ./release && \
	cp ./linux/utdf /usr/local/bin

install-mac:
	cd ./release && \
	cp ./mac/utdf /usr/local/bin
