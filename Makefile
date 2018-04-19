default: build

build:
	GOARCH=amd64 GOOS=windows go build
	upx windows-dns-test.exe

clean:
	-rm -f windows-dns-test.exe
