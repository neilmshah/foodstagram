
clean:
	rm -rf node_modules
	rm -rf package-lock.json

install:
	npm install

run: install
	node app.js

docker-build:
	docker build -t amit44/frontend .

docker-run:
	docker run -td  -p 3000:3000  amit44/frontend:latest