swagger.validate:
	swagger validate pkg/swagger/swagger.yml

swagger.doc:
	docker run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html

build:
	go build -o bin/go-swagger internal/main.go
