export tag=v1.0
root:
	export ROOT=github.com/cncamp/golang

release:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

container: release
	echo "building httpserver container"
	docker build -t cncamp/httpserver:${tag} .

container: push
	echo "pushing cncamp/httpserver"
	docker push cncamp/httpserver:v1.0
