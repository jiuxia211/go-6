

build-gateway:
	cd gateway/cmd && go build main.go
build-user:
	cd user/cmd && go build main.go
build-task:
	cd task/cmd && go build main.go