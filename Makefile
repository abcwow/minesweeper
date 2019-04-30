build:
	go build
cross_build:
	CC=/usr/x86_64-w64-mingw32/bin/gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build
	#tar -czf - MineSweeper | ssh win2 "cd $(WINHOME)/work/pk2019w18/MineSweeper; tar -xzf -; ./MineSweeper"

