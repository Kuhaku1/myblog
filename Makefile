PROJECT	=	myblog
PROJECT	=	myblog

VERSION	=	1.1.0

build:
	@echo ====================build====================
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

package:build
	@echo ====================package====================
	docker build -t $(PROJECT):$(VERSION) .

clean:
	rm -r ./myblog

package-clean: clean
	docker images | grep -E "($(PROJECT)\s+)" | awk '{print $$3}' | uniq | xargs -I {} docker rmi --force {}