.PHONY: create

create:
	cd ./utdf && \
	GOOS=windows go build -o ../release/windows/utdf.exe && \
	GOOS=linux go build -o ../release/linux/utdf && \
	go build -o ../release/mac/utdf

install-linux-mac:
	cd ./release && \
	cp utdf /usr/local/bin
