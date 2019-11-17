
all: clean

clean: 
	find . -name 'go-login' -type f -exec rm -f {} \; 
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt go-login

install:
	go install go-login

build:
	go build go-login

start:
	./go-login 

test-ping:
	curl localhost:3000/ping

test-gumball:
	curl localhost:3000/gumball

docker-build: 
	docker build -t go-login .
	docker images

docker-run:
	docker run --name go-login -td -p 8080:8080 go-login
	docker ps

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it go-login bash 

docker-clean:
	docker stop go-login
	docker rm go-login
	docker rmi go-login

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

