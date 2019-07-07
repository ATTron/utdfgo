.PHONY: create

create:
	GOOS=windows go build -o ./release/utdf.exe
	GOOS=linux go build -o ./release/utdf

install-linux-mac:
	cd $$GOPATH/src/github/ATTron/utdfgo/utdf && \
	GOBIN=/usr/local/bin/ go install
