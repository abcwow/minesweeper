#putty config:Connection->SSH->Remote command->WINHOST=win2 WINHOME=xxxx zsh
#android: 手机版 540x960

WINHOST ?= WINHOST

c:
	git add .
	git commit -v

build:
	go build main.go
cross_build:
	PATH=/usr/x86_64-w64-mingw32/bin:$(PATH) CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o main.exe
	#scp main.exe $(WINHOST):$(WINHOME)/work/git/getdedao/
	tar -czf - main.exe | ssh win2 "cd $(WINHOME)/work/git/getdedao; tar -xzf -; ./main.exe"


remote_sync:
	-git add .
	-git commit -v
	-git push
	ssh $(WINHOST) $(WINHOME)/work/git/getdedao/tools/build.sh
remote_build:
	ssh $(WINHOST) $(WINHOME)/work/git/getdedao/tools/build.sh
remote_run:
	ssh $(WINHOST) $(WINHOME)/work/git/getdedao/tools/run.sh
