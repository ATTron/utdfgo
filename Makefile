.PHONY: create

create:
	cd $$GOPATH/src/github/ATTron/utdfgo/utdf && \
	GOOS=windows go build -o ./release/utdf.exe && \
	GOOS=linux go build -o ./release/utdf

install-linux-mac:
	cd $$GOPATH/src/github/ATTron/utdfgo/release && \
	cp utdf /usr/local/bin
