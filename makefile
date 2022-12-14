openwrt:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=0   go build -o v2ray -trimpath -ldflags "-s -w -buildid=" ./main

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0   go build -o v2ray -trimpath -ldflags "-s -w -buildid=" ./main

arm:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0   go build -o v2ray -trimpath -ldflags "-s -w -buildid=" ./main

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0   go build -o v2ray -trimpath -ldflags "-s -w -buildid=" ./main

osx:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0   go build -o v2ray -trimpath -ldflags "-s -w -buildid=" ./main

